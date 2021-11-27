package repository

import (
	"github.com/yangsen996/Gin-demo/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	FindUserById(userId string) entity.User
}

type UserConnection struct {
	conn *gorm.DB
}

func NewUserConnection(db *gorm.DB) UserRepository {
	return &UserConnection{
		conn: db,
	}
}

func hashAndSalf(pwd []byte) string {
	hashPwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hashPwd)
}
func (db *UserConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalf([]byte(user.Password))
	db.conn.Save(&user)
	return user
}
func (db *UserConnection) UpdateUser(user entity.User) entity.User {
	user.Password = hashAndSalf([]byte(user.Password))
	db.conn.Save(&user)
	return user
}
func (db *UserConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.conn.Where("email =?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}
func (db *UserConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.conn.Where("email=?", email).Take(&user)
}
func (db *UserConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.conn.Where("email=?", email).Take(&user)
	return user
}
func (db *UserConnection) FindUserById(userId string) entity.User {
	var user entity.User
	db.conn.Find(&user, userId)
	return user
}
