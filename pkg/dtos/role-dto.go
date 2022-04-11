package dtos

import "github.com/gofrs/uuid"

type RoleListDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type RoleCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type RoleUpdateDTO struct {
	ID   uuid.UUID `json:"id" form:"id" binding:"required"`
	Name string    `json:"name" form:"name" binding:"required"`
}
