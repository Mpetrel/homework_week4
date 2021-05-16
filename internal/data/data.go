package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"homework_week4/internal/conf"
)

var _db *gorm.DB

type Data struct {
	db *gorm.DB
}

func NewData(c *conf.Data) (*Data, error) {
	if _db == nil {
		var err error
		_db, err = gorm.Open(sqlite.Open(c.Database.Source), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}
	return &Data{db: _db}, nil
}