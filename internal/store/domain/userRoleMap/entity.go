package userrolemap

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/role"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/user"
	"gorm.io/gorm"
)

type UserRoleMap struct {
	gorm.Model
	UserID   uint      `json:"user_id" gorm:"type:integer;not null"`
	User     user.User `json:"user" gorm:"foreignkey:UserID"`
	RoleID   uint      `json:"role_id" gorm:"type:integer;not null"`
	Role     role.Role `json:"role" gorm:"foreignkey:RoleID"`
	IsActive bool      `json:"is_active" gorm:"type:boolean;not null"`
}

func (UserRoleMap) TableName() string {
	return "user_role_map"
}
