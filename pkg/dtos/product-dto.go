package dtos

import "github.com/gofrs/uuid"

type ProductListDTO struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	SKU              string    `json:"sku"`
	ShortDescription string    `json:"shortDescription"`
	Description      string    `json:"description"`
	Price            float64   `json:"price"`
	UnitOfStock      int       `json:"unitOfStock"`
}

type ProductCreateDTO struct {
	Name             string  `json:"name" form:"name" binding:"required"`
	SKU              string  `json:"sku" form:"sku" binding:"required"`
	ShortDescription string  `json:"shortDescription" form:"shortDescription" binding:"required"`
	Description      string  `json:"description" form:"description" binding:"required"`
	Price            float64 `json:"price" form:"price" binding:"required"`
	UnitOfStock      int     `json:"unitOfStock" form:"unitOfStock" binding:"required"`
}

type ProductUpdateDTO struct {
	ID               uuid.UUID `json:"id" form:"id" binding:"required"`
	Name             string    `json:"name" form:"name" binding:"required"`
	SKU              string    `json:"sku" form:"sku" binding:"required"`
	ShortDescription string    `json:"shortDescription" form:"shortDescription" binding:"required"`
	Description      string    `json:"description" form:"description" binding:"required"`
	Price            float64   `json:"price" form:"price" binding:"required"`
	UnitOfStock      int       `json:"unitOfStock" form:"unitOfStock" binding:"required"`
}
