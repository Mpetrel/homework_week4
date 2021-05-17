package biz

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Sex string
	Age int32
}

type UserRepo interface {
	CreateUser(user *User) error
	FetchAll() ([] User, error)
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (uc *UserUsecase) Create(u *User) error {
	return uc.repo.CreateUser(u)
}

func (uc *UserUsecase) Fetch() ([] User, error) {
	return uc.repo.FetchAll()
}
