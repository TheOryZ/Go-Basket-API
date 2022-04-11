package dtos

//RegisterDTO is a struct for registering a new user
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min=3"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min=6"`
}
