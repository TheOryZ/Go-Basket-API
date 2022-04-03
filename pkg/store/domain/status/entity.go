package status

import (
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(100);not null"`
}

func (Status) TableName() string {
	return "status"
}
