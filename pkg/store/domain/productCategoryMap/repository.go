package productcategorymap

import "gorm.io/gorm"

type ProductCategoryMapRepository struct {
	db *gorm.DB
}

func NewProductCategoryMapRepository(db *gorm.DB) *ProductCategoryMapRepository {
	return &ProductCategoryMapRepository{db: db}
}
func (r *ProductCategoryMapRepository) Migration() {
	r.db.AutoMigrate(&ProductCategoryMap{})
}
