package dtos

import "github.com/gofrs/uuid"

//ProductListDTO is a struct for listing products
type ProductListDTO struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	SKU              string    `json:"sku"`
	ShortDescription string    `json:"shortDescription"`
	Description      string    `json:"description"`
	Price            float64   `json:"price"`
	UnitOfStock      uint      `json:"unitOfStock"`
}

//ProductCreateDTO is a struct for creating a new product
type ProductCreateDTO struct {
	Name             string  `json:"name" form:"name" binding:"required"`
	SKU              string  `json:"sku" form:"sku" binding:"required"`
	ShortDescription string  `json:"shortDescription" form:"shortDescription" binding:"required"`
	Description      string  `json:"description" form:"description" binding:"required"`
	Price            float64 `json:"price" form:"price" binding:"required"`
	UnitOfStock      uint    `json:"unitOfStock" form:"unitOfStock" binding:"required"`
}

//ProductUpdateDTO is a struct for updating a product
type ProductUpdateDTO struct {
	ID               uuid.UUID `json:"id" form:"id" binding:"required"`
	Name             string    `json:"name" form:"name" binding:"required"`
	SKU              string    `json:"sku" form:"sku" binding:"required"`
	ShortDescription string    `json:"shortDescription" form:"shortDescription" binding:"required"`
	Description      string    `json:"description" form:"description" binding:"required"`
	Price            float64   `json:"price" form:"price" binding:"required"`
	UnitOfStock      uint      `json:"unitOfStock" form:"unitOfStock" binding:"required"`
}

//ProductWithCategoriesDTO is a struct for listing products with categories
type ProductWithCategoriesDTO struct {
	ID               uuid.UUID         `json:"id"`
	Name             string            `json:"name"`
	SKU              string            `json:"sku"`
	ShortDescription string            `json:"shortDescription"`
	Description      string            `json:"description"`
	Price            float64           `json:"price"`
	UnitOfStock      uint              `json:"unitOfStock"`
	Categories       []CategoryListDTO `json:"categories"`
}
