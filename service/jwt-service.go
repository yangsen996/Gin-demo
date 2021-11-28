package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type JwtCustomClaim struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	SecretKey string
	issuer    string
}

func NewJwtService() JWTService {
	return &jwtService{
		issuer:    "yangsen",
		SecretKey: getSecret(),
	}
}
func getSecret() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "yangsen"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(userID string) string {
	claims := &JwtCustomClaim{
		userID,
		jwt.StandardClaims{
			Issuer:    j.issuer,
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	t, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.SecretKey), nil
	})
}
