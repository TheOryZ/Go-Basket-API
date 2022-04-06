package userrolemap

import "gorm.io/gorm"

type UserRoleMapRepository struct {
	db *gorm.DB
}

func NewUserRoleMapRepository(db *gorm.DB) *UserRoleMapRepository {
	return &UserRoleMapRepository{db: db}
}
func (r *UserRoleMapRepository) Migration() {
	r.db.AutoMigrate(&UserRoleMap{})
}
