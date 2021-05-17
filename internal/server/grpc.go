package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	v1 "homework_week4/api/user/v1"
	"homework_week4/internal/conf"
	"homework_week4/internal/service"
)

func NewGRPCServer(c *conf.Server, user *service.UserService) *grpc.Server {
	srv := grpc.NewServer()
	v1.RegisterUserServer(srv, user)
	reflection.Register(srv)
	return srv
}
