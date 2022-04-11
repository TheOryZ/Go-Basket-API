package dtos

import "github.com/gofrs/uuid"

//StatusListDTO is a struct for listing statuses
type StatusListDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

//StatusCreateDTO is a struct for creating a new status
type StatusCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
}

//StatusUpdateDTO is a struct for updating a status
type StatusUpdateDTO struct {
	ID   uuid.UUID `json:"id" form:"id" binding:"required"`
	Name string    `json:"name" form:"name" binding:"required"`
}
