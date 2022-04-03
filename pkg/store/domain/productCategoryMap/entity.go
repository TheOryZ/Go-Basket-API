package productcategorymap

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/domain/category"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/store/domain/product"
	"gorm.io/gorm"
)

type ProductCategoryMap struct {
	gorm.Model
	ProductID  uint              `json:"product_id" gorm:"type:integer;not null"`
	Product    product.Product   `json:"product" gorm:"foreignkey:ProductID"`
	CategoryID uint              `json:"category_id" gorm:"type:integer;not null"`
	Category   category.Category `json:"category" gorm:"foreignkey:CategoryID"`
	IsActive   bool              `json:"is_active" gorm:"type:boolean;not null"`
}

func (productcategorymap ProductCategoryMap) TableName() string {
	return "product_category_map"
}
