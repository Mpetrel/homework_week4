package data

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)


var ProviderSet = wire.NewSet(NewData, NewUserRepo)

type Data struct {
	db *gorm.DB
}

func NewData() (*Data, error) {
	return &Data{db: db}, nil
}