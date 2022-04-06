package role

import "gorm.io/gorm"

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}
func (r *RoleRepository) Migration() {
	r.db.AutoMigrate(&Role{})
}
