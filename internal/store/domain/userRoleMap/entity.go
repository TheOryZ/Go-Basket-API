package userrolemap

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/role"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/user"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserRoleMap struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User     user.User `json:"user" gorm:"foreignkey:UserID"`
	RoleID   uuid.UUID `json:"role_id" gorm:"type:uuid;not null"`
	Role     role.Role `json:"role" gorm:"foreignkey:RoleID"`
	IsActive bool      `json:"is_active" gorm:"type:boolean;not null"`
}

func (UserRoleMap) TableName() string {
	return "user_role_map"
}
