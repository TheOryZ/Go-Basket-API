package order

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/product"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/status"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/user"
	"github.com/gofrs/uuid"
)

type Order struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID    uuid.UUID       `json:"user_id" gorm:"type:uuid;not null"`
	User      user.User       `json:"user" gorm:"foreignkey:UserID"`
	ProductID uuid.UUID       `json:"product_id" gorm:"type:uuid;not null"`
	Product   product.Product `json:"product" gorm:"foreignkey:ProductID"`
	Price     float64         `json:"price" gorm:"type:float;not null"`
	StatusID  uuid.UUID       `json:"status_id" gorm:"type:uuid;not null"`
	Status    status.Status   `json:"status" gorm:"foreignkey:StatusID"`
	CreatedAt string          `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt string          `json:"updated_at" gorm:"type:timestamp;not null"`
	DeletedAt string          `json:"deleted_at" gorm:"type:timestamp;default:null"`
	IsActive  bool            `json:"is_active" gorm:"type:boolean;not null"`
}

func (Order) TableName() string {
	return "orders"
}
