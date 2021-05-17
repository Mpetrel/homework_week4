package service

import (
	"context"
	"github.com/pkg/errors"
	v1 "homework_week4/api/user/v1"
	"homework_week4/internal/biz"
	"time"
)

// UserService user info operate
type UserService struct {
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{
		uc: uc,
	}
}

// AddUser implements grpc user server
func (us *UserService) AddUser(context context.Context, in *v1.AddUserRequest) (*v1.AddUserReply, error) {
	user := biz.User{
		Name:  in.GetName(),
		Sex:   in.GetSex(),
		Age:   in.GetAge(),
	}
	user.ID = uint(time.Now().Nanosecond())
	err := us.uc.Create(&user)
	if err != nil {
		return nil, errors.Wrap(err, "create user failed")
	}
	reply := &v1.AddUserReply{
		Code:    0,
		Message: "OK",
	}
	return reply, nil
}

func (us *UserService) FetchUser(context.Context, *v1.EmptyMessage) (*v1.FetchUsersReply, error) {
	users, err := us.uc.Fetch()
	if err != nil {
		return nil, errors.Wrap(err, "fetch user failed")
	}
	var replyUsers [] *v1.FetchUsersReply_UserInfo
	for _, user := range users {
		replyUsers = append(replyUsers, &v1.FetchUsersReply_UserInfo{
			Id: int32(user.ID),
			Name: user.Name,
			Sex:  user.Sex,
			Age:  user.Age,
		})
	}
	reply := &v1.FetchUsersReply{Users: replyUsers}
	return reply, nil
}


