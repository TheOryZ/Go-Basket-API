package productcategorymap

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/category"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/product"
	"github.com/gofrs/uuid"
)

type ProductCategoryMap struct {
	ID         uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4()"`
	ProductID  uuid.UUID         `json:"product_id" gorm:"type:uuid;not null"`
	Product    product.Product   `json:"product" gorm:"foreignkey:ProductID"`
	CategoryID uuid.UUID         `json:"category_id" gorm:"type:uuid;not null"`
	Category   category.Category `json:"category" gorm:"foreignkey:CategoryID"`
	CreatedAt  string            `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt  string            `json:"updated_at" gorm:"type:timestamp;not null"`
	DeletedAt  string            `json:"deleted_at" gorm:"type:timestamp;default:null"`
	IsActive   bool              `json:"is_active" gorm:"type:boolean;not null"`
}

func (productcategorymap ProductCategoryMap) TableName() string {
	return "product_category_map"
}
