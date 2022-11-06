SERVICE_TYPE?=service
PROJECT_NAME?=github.com/weblfe/flyfire
APP_LAYOUT?=https://github.com/go-kratos/beer-shop
SERVICE_LAYOUT?=https://gitee.com/go-kratos/kratos-layout
COMMIT?=update
RELEASE?=$(date)
BRANCH?=release/sit/$(RELEASE)

.PHONY: init
# init env
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
	go get -u github.com/google/wire/cmd/wire
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2
	go get -u github.com/envoyproxy/protoc-gen-validate

.PHONY: dos2unix
# 批量转换所有文件格式为LF
dos2unix:
	find . -type f | xargs dos2unix


.PHONY: merge merge-without-commit
# 合并
merge:
	git add .
	git commit -m '$(COMMIT)'
	git push origin $(subst * ,,$(shell git branch | grep '*'))
	make merge-without-commit

merge-without-commit:
	git checkout $(BRANCH)
	git pull origin $(BRANCH)
	git merge $(subst * ,,$(shell git branch | grep '*')) --no-commit --no-ff
	make api
	make proto
	git add .
	git commit -m 'merge from $(subst * ,,$(shell git branch | grep '*'))'
	git push origin $(BRANCH)
	git checkout $(subst * ,,$(shell git branch | grep '*'))

.PHONY: fetch-unrelease
# 拉取
fetch-unrelease:
	git pull origin $(subst * ,,$(shell git branch | grep '*')) --no-commit --no-ff
	cat unrelease.txt | xargs -I % git merge % --no-commit --no-ff

.PHONY: api
# generate api
api:
	find app -mindepth 2 -maxdepth 2 -type d -print | grep -v sql | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) api'

.PHONY: wire
# generate wire
wire:
	find app -mindepth 2 -maxdepth 2 -type d -print | grep -v sql | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) wire'

.PHONY: proto
# generate proto
proto:
	find app -mindepth 2 -maxdepth 2 -type d -print | grep -v sql | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) proto'

.PHONY: build
# generate build
build:
	find app -mindepth 2 -maxdepth 2 -type d -print | grep -v sql | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) build'

.PHONY: app
# generate app
app:
	rm -rf app/$(APPNAME)
	mkdir -p app/$(APPNAME)
	kratos new $(PROJECT_NAME)/app/$(APPNAME)/$(SERVICE_TYPE) -r $(APP_LAYOUT) -b $
	(SERVICE_TYPE)
	rm -f $(SERVICE_TYPE)/go.{mod,sum}
	cp -r $(SERVICE_TYPE) app/$(APPNAME)/
	rm -rf $(SERVICE_TYPE)
	mkdir -p api/$(APPNAME)/$(SERVICE_TYPE)/v1
	touch api/$(APPNAME)/$(SERVICE_TYPE)/v1/$(APPNAME).proto

.PHONY: service
# generate service, eg: make service APPNAME=yourServerName PROJECT_NAME=yourProjectName
service:
	rm -rf app/$(APPNAME)
	mkdir -p app/$(APPNAME)
	kratos new $(PROJECT_NAME)/app/$(APPNAME)/$(SERVICE_TYPE) -r $(SERVICE_LAYOUT)
	@rm -rf $(SERVICE_TYPE)/go.{mod,sum} $(SERVICE_TYPE)/.gitignore $(SERVICE_TYPE)/LICENSE $(SERVICE_TYPE)/openapi.yaml \
     $(SERVICE_TYPE)/third_party  $(SERVICE_TYPE)/api
	@echo "include ../../../app_makefile" > $(SERVICE_TYPE)/Makefile
	@cp -r $(SERVICE_TYPE) app/$(APPNAME)/
	@rm -rf $(SERVICE_TYPE)
	@mkdir -p app/$(APPNAME)/sql
	@rm -rf api/$(APPNAME)/$(SERVICE_TYPE)/v1/
	@mkdir -p api/$(APPNAME)/$(SERVICE_TYPE)/v1
	@touch api/$(APPNAME)/$(SERVICE_TYPE)/v1/$(APPNAME).proto
	@echo "\
    syntax = \"proto3\";\n\
    package api.$(APPNAME).service.v1; \n\
    option go_package = \"$(PROJECT_NAME)/api/$(APPNAME)/service/v1;v1\";\n\
    option java_multiple_files = true; \n\
    option java_package = \"api.$(APPNAME).service.v1\";\n" >> api/$(APPNAME)/$(SERVICE_TYPE)/v1/$(APPNAME).proto