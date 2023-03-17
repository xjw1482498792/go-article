package common

import (
	"goo/model"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	args := "root:123456@(localhost)/ginessential?charset=utf8mb4&parseTime=True&loc=Local"
	driverName := "mysql"
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db

}

func GetDB() *gorm.DB {

	return DB
}
