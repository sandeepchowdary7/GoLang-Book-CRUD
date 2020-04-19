package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:book@/book?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect Database")
	}

	db.AutoMigrate(&Book{})

	return db
}