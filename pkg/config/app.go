package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	database, err := gorm.Open("mysql", "root:topsecret@/bookstore?charset=utf8&parseTime=True")

	if err != nil {
		panic(err)
	}

	db = database
}

func GetDB() *gorm.DB {
	return db
}
