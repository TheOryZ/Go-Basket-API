package product

import (
	"github.com/gofrs/uuid"
)

type Product struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name             string    `json:"name" gorm:"type:varchar(255);not null"`
	SKU              string    `json:"sku" gorm:"type:varchar(255);not null"`
	ShortDescription string    `json:"short_description" gorm:"type:varchar(255);not null"`
	Description      string    `json:"description" gorm:"type:varchar(255);not null"`
	Price            float64   `json:"price" gorm:"type:float;not null"`
	UnitOfStock      int       `json:"unit_of_stock" gorm:"type:integer;not null"`
	CreatedAt        string    `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt        string    `json:"updated_at" gorm:"type:timestamp;not null"`
	DeletedAt        string    `json:"deleted_at" gorm:"type:timestamp;default:null"`
	IsActive         bool      `json:"is_active" gorm:"type:boolean;not null"`
}

func (Product) TableName() string {
	return "products"
}
