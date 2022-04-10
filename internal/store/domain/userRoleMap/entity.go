package userrolemap

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/role"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/user"
	"github.com/gofrs/uuid"
)

type UserRoleMap struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User      user.User `json:"user" gorm:"foreignkey:UserID"`
	RoleID    uuid.UUID `json:"role_id" gorm:"type:uuid;not null"`
	Role      role.Role `json:"role" gorm:"foreignkey:RoleID"`
	CreatedAt string    `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt string    `json:"updated_at" gorm:"type:timestamp;not null"`
	DeletedAt string    `json:"deleted_at" gorm:"type:timestamp;default:null"`
	IsActive  bool      `json:"is_active" gorm:"type:boolean;not null"`
}

func (UserRoleMap) TableName() string {
	return "user_role_map"
}
