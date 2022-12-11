package router

import (
	"gin_project/controller"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/user/auth/register", controller.Register)
	r.POST("/user/auth/login", controller.Login)
	return r
}
