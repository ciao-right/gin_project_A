package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var r = gin.Default()

func main() {
	r.POST("/user/auth/register", func(ctx *gin.Context) {
		//获取传值
		userName := ctx.Query("userName")
		password := ctx.Query("password")
		tel := ctx.Query("tel")
		// 判断是否正确
		if len(password) < 6 {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "请输入大于六位的密码",
				"code": 200,
				"data": nil,
			})
			return
		}
		if len(tel) != 11 {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "请输入正确的手机号",
				"code": 200,
				"data": nil,
			})
			return

		}
		if isEmptyString(userName) {
			userName = RandString(10)
		}
		log.Println(userName, password, tel)
		//返回

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
