package role

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name" gorm:"type:varchar(255);not null"`
}

func (Role) TableName() string {
	return "roles"
}
