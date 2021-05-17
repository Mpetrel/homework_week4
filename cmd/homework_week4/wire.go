// +build wireinject

package main

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	"homework_week4/internal/biz"
	"homework_week4/internal/conf"
	"homework_week4/internal/data"
	"homework_week4/internal/server"
	"homework_week4/internal/service"
)

func initServer(*conf.Server, *conf.Data) (*grpc.Server, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet))
}