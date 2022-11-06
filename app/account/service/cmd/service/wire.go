//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"githu.com/weblfe/flyfire/app/account/service/internal/biz"
	"githu.com/weblfe/flyfire/app/account/service/internal/conf"
	"githu.com/weblfe/flyfire/app/account/service/internal/data"
	"githu.com/weblfe/flyfire/app/account/service/internal/server"
	"githu.com/weblfe/flyfire/app/account/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
