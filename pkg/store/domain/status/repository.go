package status

import "gorm.io/gorm"

type StatusRepository struct {
	db *gorm.DB
}

func NewStatusRepository(db *gorm.DB) *StatusRepository {
	return &StatusRepository{db: db}
}
func (r *StatusRepository) Migration() {
	r.db.AutoMigrate(&Status{})
}
