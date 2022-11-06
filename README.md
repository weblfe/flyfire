# FlyFire  [飞火---星星之火]

## Intro [项目简介]

```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

## Project Layout [项目结构]

```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Start  [项目开始/项目部署]


### make build [编译部署]
```bash
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```

### docker run [docker部署]
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

## DependOn [项目依赖]


## Roadmap [项目规划]


