package middleware

import (
	"log"
	"net/http"

	"github.com/yangsen996/Gin-demo/comm"
	"github.com/yangsen996/Gin-demo/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeHandler(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("authorization")
		if authHeader == "" {
			response := comm.BuildErrorResponse("failed to process request", "token not found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("claims[user_id]", claims["user_id"])
			log.Println("claims[issuer]", claims["issuer"])
		} else {
			response := comm.BuildErrorResponse("token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
