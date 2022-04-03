package category

import "gorm.io/gorm"

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}
func (r *CategoryRepository) Migration() {
	r.db.AutoMigrate(&Category{})
}
