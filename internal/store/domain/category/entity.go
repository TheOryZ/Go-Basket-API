package category

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	IsActive bool   `json:"is_active" gorm:"type:boolean;not null"`
}

func (Category) TableName() string {
	return "categories"
}
