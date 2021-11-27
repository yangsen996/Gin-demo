package main

import (
	"github.com/yangsen996/Gin-demo/config"
	"github.com/yangsen996/Gin-demo/controller"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var (
	authController controller.AuthController = controller.NewAuthController()
	db             *gorm.DB
)

func main() {
	//初始化数据库
	db = config.ConnDB()
	defer config.Close(db)
	//初始化路由
	router := gin.Default()
	authRouter := router.Group("api/auth")
	{
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/register", authController.Register)
	}
	router.Run(":9999")
}
