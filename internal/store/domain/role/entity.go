package role

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(255);not null"`
}

func (Role) TableName() string {
	return "roles"
}
