package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yangsen996/Gin-demo/comm"
	"github.com/yangsen996/Gin-demo/dto"
	"github.com/yangsen996/Gin-demo/entity"
	"github.com/yangsen996/Gin-demo/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (a *authController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDTO
	err := ctx.ShouldBind(&loginDto)
	if err != nil {
		comm.BuildErrorResponse("参数错误", err.Error(), comm.EmptyObj{})
	}
	authRes := a.authService.VerifyCredential(loginDto.Email, loginDto.Password)
	if v, ok := authRes.(entity.User); ok {
		genToken := a.jwtService.GenerateToken(strconv.FormatUint(v.Id, 10))
		v.Token = genToken
		response := comm.BuildResponse(true, "success", v)
		ctx.JSON(http.StatusOK, response)
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, "请注册")
}
func (a *authController) Register(ctx *gin.Context) {
	var register dto.RegisterDTO
	err := ctx.ShouldBind(&register)
	if err != nil {
		response := comm.BuildErrorResponse("参数错误", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	// 验证数据库存在不存在用户
	if !a.authService.IsDuplicateEmail(register.Email) {
		ctx.JSON(http.StatusBadRequest, "邮箱已存在")
	} else {
		addUser := a.authService.CreateUser(register)
		toekn := a.jwtService.GenerateToken(strconv.FormatUint(addUser.Id, 10))
		addUser.Token = toekn
		response := comm.BuildResponse(true, "success", addUser)
		ctx.JSON(http.StatusOK, response)
	}
}
