package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"homework_week4/internal/biz"
	"homework_week4/internal/conf"
)

var db *gorm.DB

func SetupDB(c *conf.Data) error {
	var err error
	db, err = gorm.Open(sqlite.Open(c.Database.Source), &gorm.Config{})
	checkTable()
	return err
}

func checkTable() {
	_ = db.AutoMigrate(&biz.User{})
}