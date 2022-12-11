package controller

import (
	"gin_project/common"
	"gin_project/model"
	"gin_project/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Register(ctx *gin.Context) {
	//获取传值
	db := common.GetDb()
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

	if utils.IsEmptyString(userName) {
		userName = utils.RandString(10)
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
	//返回
	newUser := model.User{
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
}

func Login(ctx *gin.Context) {
	//获取参数
	userName := ctx.PostForm("userName")
	password := ctx.PostForm("password")
	//判断有没有输入参数
	if utils.IsEmptyString(userName) || utils.IsEmptyString(password) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "请输入用户名或者密码",
			"code": 201,
			"data": 0,
		})
		return
	}
	//判断是否输入正确
	var user model.User
	common.GetDb().Where("name = ?", userName).First(&user)
	if user.ID != 0 {
		if user.Password == password {
			ctx.JSON(http.StatusOK, gin.H{
				"code": "200",
				"msg":  "success",
				"data": 1,
			})
		}
	}
}

func isTelExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("tel = ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
