package dtos

import "github.com/gofrs/uuid"

type StatusListDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type StatusCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type StatusUpdateDTO struct {
	ID   uuid.UUID `json:"id" form:"id" binding:"required"`
	Name string    `json:"name" form:"name" binding:"required"`
}
