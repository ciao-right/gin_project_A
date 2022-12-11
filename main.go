package main

import (
	"gin_project/router/user"
	"github.com/gin-gonic/gin"
)

func main() {
	var r = gin.Default()
	router.CollectRouter(r)
	r.Run(":3000")
}
