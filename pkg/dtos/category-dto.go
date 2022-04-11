package dtos

import "github.com/gofrs/uuid"

type CategoryListDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CategoryCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type CategoryUpdateDTO struct {
	ID   uuid.UUID `json:"id" form:"id" binding:"required"`
	Name string    `json:"name" form:"name" binding:"required"`
}
