package data

import "homework_week4/internal/biz"

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (r *userRepo) CreateUser(user *biz.User) error {
	return r.data.db.Create(user).Error
}

func (r *userRepo) FetchAll() (users [] biz.User, err error) {
	err = r.data.db.Find(&users).Error
	return
}