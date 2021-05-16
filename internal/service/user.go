package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "homework_week4/api/user/v1"
)

// UserService user info operate
type UserService struct {
}

// AddUser implements grpc user server
func (us *UserService) AddUser(ctx context.Context, in *v1.AddUserRequest) (*v1.AddUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}