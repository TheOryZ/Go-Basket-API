package status

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name" gorm:"type:varchar(100);not null"`
}

func (Status) TableName() string {
	return "status"
}
