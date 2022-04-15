package dtos

import "github.com/gofrs/uuid"

//CategoryListDTO is a struct for listing categories
type CategoryListDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

//CategoryCreateDTO is a struct for creating a new category
type CategoryCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
}

//CategoryUpdateDTO is a struct for updating a category
type CategoryUpdateDTO struct {
	ID   uuid.UUID `json:"id" form:"id" binding:"required"`
	Name string    `json:"name" form:"name" binding:"required"`
}

//CategoryWithProductsDTO is a struct for listing categories with products
type CategoryWithProductsDTO struct {
	ID       uuid.UUID        `json:"id"`
	Name     string           `json:"name"`
	Products []ProductListDTO `json:"products"`
}
