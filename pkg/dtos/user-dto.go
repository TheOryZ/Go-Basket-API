package dtos

import "github.com/gofrs/uuid"

type UserListDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type UserCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserUpdateDTO struct {
	ID       uuid.UUID `json:"id" form:"id" binding:"required"`
	Name     string    `json:"name" form:"name" binding:"required"`
	Email    string    `json:"email" form:"email" binding:"required"`
	Password string    `json:"password" form:"password" binding:"required"`
}
