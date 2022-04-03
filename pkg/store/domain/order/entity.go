package order

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/domain/product"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/domain/status"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/domain/user"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID    uint            `json:"user_id" gorm:"type:int;not null"`
	User      user.User       `json:"user" gorm:"foreignkey:UserID"`
	ProductID uint            `json:"product_id" gorm:"type:int;not null"`
	Product   product.Product `json:"product" gorm:"foreignkey:ProductID"`
	Price     float64         `json:"price" gorm:"type:float;not null"`
	StatusID  uint            `json:"status_id" gorm:"type:int;not null"`
	Status    status.Status   `json:"status" gorm:"foreignkey:StatusID"`
	IsActive  bool            `json:"is_active" gorm:"type:boolean;not null"`
}

func (Order) TableName() string {
	return "orders"
}
