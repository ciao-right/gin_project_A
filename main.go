package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var r = gin.Default()

func main() {
	db := InitDatabase()

	r.POST("/user/auth/register", func(ctx *gin.Context) {
		//获取传值
		userName := ctx.PostForm("userName")
		password := ctx.PostForm("password")
		tel := ctx.PostForm("tel")
		if isTelExist(db, tel) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "该手机号已经存在了",
				"code": 200,
				"data": 0,
			})
			return
		}
		// 判断是否正确

		if isEmptyString(userName) {
			userName = RandString(10)
		}

		if len(tel) != 11 {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "请输入正确的手机号",
				"code": 200,
				"data": nil,
			})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "请输入大于六位的密码",
				"code": 200,
				"data": password,
			})
			return
		}
		log.Println(userName, password, tel)
		//返回
		newUser := User{
			Name:     userName,
			Password: password,
			Tel:      tel,
		}
		db.Create(&newUser)
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "注册成功",
			"code": 200,
			"data": 1,
		})
	})
	r.Run(":3000")

}

func isEmptyString(str string) bool {
	return len(str) == 0
}
func RandString(num int) string {
	rand.Seed(time.Now().Unix())
	letters := []byte("asdfasdhajksbhvajdsifuaiosdfaksndfkl")
	res := make([]byte, num)
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}
	return string(res)
}

func InitDatabase() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "gin_db"
	username := "root"
	password := "1115774750"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Tel      string
	Password string
}

func isTelExist(db *gorm.DB, tel string) bool {
	var user User
	db.Where("tel = ?", tel).First(&user)
	log.Println(user)
	if user.ID != 0 {
		return true
	}
	return false
}
