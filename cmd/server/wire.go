//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/ChungKiet/cake-interview/internal/biz"
	"github.com/ChungKiet/cake-interview/internal/conf"
	data "github.com/ChungKiet/cake-interview/internal/data"
	"github.com/ChungKiet/cake-interview/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	coreConf "github.com/indikay/go-core/conf"
	"github.com/indikay/go-core/server"
	coreService "github.com/indikay/go-core/service"
)

// initApp init kratos application.
func initApp(*coreConf.Server, *conf.Data, log.Logger) (coreService.Service, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, initService))
}
