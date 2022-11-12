//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
		"github.com/weblfe/flyfire/app/account/service/internal/biz"
		"github.com/weblfe/flyfire/app/account/service/internal/conf"
		"github.com/weblfe/flyfire/app/account/service/internal/data"
		"github.com/weblfe/flyfire/app/account/service/internal/server"
		"github.com/weblfe/flyfire/app/account/service/internal/service"

		"github.com/go-kratos/kratos/v2"
		"github.com/go-kratos/kratos/v2/log"
		"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp([]string,*conf.Registry,*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
