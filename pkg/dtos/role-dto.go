package dtos

import "github.com/gofrs/uuid"

//RoleListDTO is a struct for listing roles
type RoleListDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

//RoleCreateDTO is a struct for creating a new role
type RoleCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
}

//RoleUpdateDTO is a struct for updating a role
type RoleUpdateDTO struct {
	ID   uuid.UUID `json:"id" form:"id" binding:"required"`
	Name string    `json:"name" form:"name" binding:"required"`
}

//RoleWithUserDTO is a struct for listing roles with users
type RoleWithUserDTO struct {
	ID   uuid.UUID     `json:"id"`
	Name string        `json:"name"`
	User []UserListDTO `json:"user"`
}
