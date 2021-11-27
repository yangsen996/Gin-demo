package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct{}

func NewAuthController() AuthController {
	return &authController{}
}

func (a *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "login"})
}
func (a *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "register"})
}
