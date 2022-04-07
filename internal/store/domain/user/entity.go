package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	IsActive bool   `json:"is_active" gorm:"type:boolean;not null"`
}

func (User) TableName() string {
	return "users"
}
