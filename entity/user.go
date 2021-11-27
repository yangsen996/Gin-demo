package entity

type User struct {
	Id       uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"uniqueIndex;type:varchar(255)"`
	Password string `json:"password" gorm:"->;<-;not null"`
	Token    string `josn:"token" gorm:"-"`
}
