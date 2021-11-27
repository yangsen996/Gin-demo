package main

import (
	"api_auth/config"
	"api_auth/controller"

	"github.com/gin-gonic/gin"
)

var authController controller.AuthController = controller.NewAuthController()

func main() {
	//初始化数据库
	config.ConnDB()
	//初始化路由
	router := gin.Default()
	authRouter := router.Group("api/auth")
	{
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/register", authController.Register)
	}
	router.Run(":9999")
}
