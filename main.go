package main

import (
	"goo/common"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(110);not null;unique"`
	Password  string `gorm:"size:255;not null`
}

func main() {

	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}
