package service

import (
	"log"

	"github.com/yangsen996/Gin-demo/dto"
	"github.com/yangsen996/Gin-demo/entity"
	"github.com/yangsen996/Gin-demo/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) bool
	FindByEmail(email string) entity.User
	CreateUser(user dto.RegisterDTO) entity.User
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func comparePassword(hashedPwd string, password []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, password)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (a *authService) VerifyCredential(email string, password string) interface{} {
	res := a.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		compare := comparePassword(v.Password, []byte(password))
		if v.Email == email && compare {
			return res
		}
		return false
	}
	return false
}
func (a *authService) IsDuplicateEmail(email string) bool {
	res := a.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}
func (a *authService) FindByEmail(email string) entity.User {
	return a.userRepository.FindByEmail(email)
}
func (a *authService) CreateUser(userReq dto.RegisterDTO) entity.User {
	user := entity.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	res := a.userRepository.InsertUser(user)
	return res
}
