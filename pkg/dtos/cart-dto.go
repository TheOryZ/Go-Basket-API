package dtos

import "github.com/gofrs/uuid"

//CartListDTO is a struct for listing carts
type CartListDTO struct {
	ID       uuid.UUID      `json:"id"`
	User     UserListDTO    `json:"user"`
	Product  ProductListDTO `json:"product"`
	Quantity uint           `json:"quantity"`
	Price    float64        `json:"price"`
	Status   StatusListDTO  `json:"status"`
}

//CartCreateDTO is a struct for creating a new cart
type CartCreateDTO struct {
	UserID    uuid.UUID `json:"user_id" form:"user_id"`
	ProductID uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
	Quantity  uint      `json:"quantity" form:"quantity" binding:"required"`
	Price     float64   `json:"price" form:"price"`
	StatusID  uuid.UUID `json:"status_id" form:"status_id"`
}

//CartUpdateDTO is a struct for updating a cart
type CartUpdateDTO struct {
	ID        uuid.UUID `json:"id" form:"id" binding:"required"`
	UserID    uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
	ProductID uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
	Quantity  uint      `json:"quantity" form:"quantity" binding:"required"`
	StatusID  uuid.UUID `json:"status_id" form:"status_id" binding:"required"`
}
