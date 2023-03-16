package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

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
	db := InitDB()
	defer db.Close()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {

		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		}
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, telephone, password)
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
			return
		}
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		ctx.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true

	}
	return false
}

func RandomString(n int) string {
	var letters = []byte("asfsdafsdafasdfdsafslkdajKLASFJKSDJGKLASJFKLDAJKFL")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	args := "root:123456@(localhost)/ginessential?charset=utf8mb4&parseTime=True&loc=Local"
	driverName := "mysql"
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	db.AutoMigrate(&User{})
	return db

}
