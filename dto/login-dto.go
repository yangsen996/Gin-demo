package dto

type LoginDTO struct {
	Email    string `json:"email" form:"email" bindind:"required" validate:"email"`
	Password string `json:"password" form:"password" bindind:"required" validate:"min:6"`
}
