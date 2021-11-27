package entity

type Book struct {
	Id          uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string `json:"title" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:text"`
	UserId      uint64 `json:"user_id" gorm:"-"`
	User        User   `json:"user" gorm:"foreignkey:UserId;constraint:OnDelete:SET NILL"`
}
