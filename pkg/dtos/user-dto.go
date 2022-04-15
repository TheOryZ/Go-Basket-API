package dtos

import "github.com/gofrs/uuid"

//UserListDTO is a struct for listing users
type UserListDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

//UserCreateDTO is a struct for creating a new user
type UserCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min=6"`
}

//UserUpdateDTO is a struct for updating a user
type UserUpdateDTO struct {
	ID       uuid.UUID `json:"id" form:"id" binding:"required"`
	Name     string    `json:"name" form:"name" binding:"required"`
	Email    string    `json:"email" form:"email" binding:"required" validate:"email"`
	Password string    `json:"password" form:"password" binding:"required" validate:"min=6"`
}

//UserWithRolesDTO is a struct for listing users with roles
type UserWithRolesDTO struct {
	ID    uuid.UUID     `json:"id"`
	Name  string        `json:"name"`
	Email string        `json:"email"`
	Roles []RoleListDTO `json:"roles"`
}
