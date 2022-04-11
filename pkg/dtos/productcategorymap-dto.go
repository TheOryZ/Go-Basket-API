package dtos

import "github.com/gofrs/uuid"

//ProductCategoryMapListDTO is a struct for listing productcategorymaps
type ProductCategoryMapListDTO struct {
	ID       uuid.UUID        `json:"id"`
	Product  *ProductListDTO  `json:"product"`
	Category *CategoryListDTO `json:"category"`
}

//ProductCategoryMapCreateDTO is a struct for creating a new productcategorymap
type ProductCategoryMapCreateDTO struct {
	ProductID  uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
	CategoryID uuid.UUID `json:"category_id" form:"category_id" binding:"required"`
}

//ProductCategoryMapUpdateDTO is a struct for updating a productcategorymap
type ProductCategoryMapUpdateDTO struct {
	ID         uuid.UUID `json:"id" form:"id" binding:"required"`
	ProductID  uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
	CategoryID uuid.UUID `json:"category_id" form:"category_id" binding:"required"`
}
