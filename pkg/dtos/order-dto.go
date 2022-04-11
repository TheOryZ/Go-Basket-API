package dtos

import "github.com/gofrs/uuid"

//OrderListDTO is a struct for listing orders
type OrderListDTO struct {
	ID       uuid.UUID      `json:"id"`
	User     UserListDTO    `json:"user"`
	Product  ProductListDTO `json:"product"`
	Quantity int            `json:"quantity"`
	Price    float64        `json:"price"`
	Status   StatusListDTO  `json:"status"`
}

//OrderCreateDTO is a struct for creating a new order
type OrderCreateDTO struct {
	UserID    uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
	ProductID uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" form:"quantity" binding:"required"`
	Price     float64   `json:"price" form:"price" binding:"required"`
	StatusID  uuid.UUID `json:"status_id" form:"status_id" binding:"required"`
}

//OrderUpdateDTO is a struct for updating a order
type OrderUpdateDTO struct {
	ID        uuid.UUID `json:"id" form:"id" binding:"required"`
	UserID    uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
	ProductID uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" form:"quantity" binding:"required"`
	Price     float64   `json:"price" form:"price" binding:"required"`
	StatusID  uuid.UUID `json:"status_id" form:"status_id" binding:"required"`
}
