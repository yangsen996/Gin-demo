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

type userConnection struct {
	conn *gorm.DB
}

func NewUserConnection(db *gorm.DB) UserRepository {
	return &userConnection{
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
func (db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalf([]byte(user.Password))
	db.conn.Save(&user)
	return user
}
func (db *userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalf([]byte(user.Password))
	} else {
		var tmp entity.User
		db.conn.Find(&tmp, user.Id)
		user.Password = tmp.Password
	}
	db.conn.Save(&user)
	return user
}
func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.conn.Where("email =?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}
func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.conn.Where("email=?", email).Take(&user)
}
func (db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.conn.Where("email=?", email).Take(&user)
	return user
}
func (db *userConnection) FindUserById(userId string) entity.User {
	var user entity.User
	db.conn.Find(&user, userId)
	return user
}
