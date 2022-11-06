// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"githu.com/weblfe/flyfire/app/account/service/internal/biz"
	"githu.com/weblfe/flyfire/app/account/service/internal/conf"
	"githu.com/weblfe/flyfire/app/account/service/internal/data"
	"githu.com/weblfe/flyfire/app/account/service/internal/server"
	"githu.com/weblfe/flyfire/app/account/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	accountRepo := data.NewAccountRepo(dataData, logger)
	accountUseCase := biz.NewAccountUseCase(accountRepo, logger)
	accountService := service.NewAccountService(accountUseCase)
	grpcServer := server.NewGRPCServer(confServer, accountService, logger)
	httpServer := server.NewHTTPServer(confServer, accountService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}