package main

import (
	"github.com/yangsen996/Gin-demo/config"
	"github.com/yangsen996/Gin-demo/controller"
	"github.com/yangsen996/Gin-demo/repository"
	"github.com/yangsen996/Gin-demo/service"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var (
	db             *gorm.DB
	userRepository repository.UserRepository = repository.NewUserConnection(db)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	jwtService     service.JWTService        = service.NewJwtService()
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
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
