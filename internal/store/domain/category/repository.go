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

//Create a new category
func (r *CategoryRepository) Create(category *Category) error {
	return r.db.Create(category).Error
}

//Update a category
func (r *CategoryRepository) Update(category *Category) error {
	return r.db.Save(category).Error
}

//Delete a category
func (r *CategoryRepository) Delete(category *Category) error {
	return r.db.Delete(category).Error
}

//Delete a category by id
func (r *CategoryRepository) DeleteByID(id uint) error {
	category := Category{}
	category.ID = id
	return r.db.Delete(&category).Error
}

//Find all categories
func (r *CategoryRepository) FindAll() ([]Category, error) {
	var categories []Category
	err := r.db.Where("deleted_at = ?", nil).Find(&categories).Error
	return categories, err
}

//Find all with pagination  //TODO: check this
func (r *CategoryRepository) FindAllWithPagination(page int, limit int) ([]Category, error) {
	var categories []Category
	err := r.db.Where("deleted_at = ?", nil).Offset(page).Limit(limit).Find(&categories).Error
	return categories, err
}

//Count all categories
func (r *CategoryRepository) CountAll() (int64, error) {
	var count int64
	err := r.db.Model(&Category{}).Where("deleted_at = ?", nil).Count(&count).Error
	return count, err
}

//Find a category by id
func (r *CategoryRepository) FindByID(id uint) (*Category, error) {
	category := Category{}
	err := r.db.First(&category, id).Error
	return &category, err
}

//Find a category by name
func (r *CategoryRepository) FindByName(name string) (*Category, error) {
	category := Category{}
	err := r.db.Where("name = ?", name).First(&category).Error
	return &category, err
}

//Search categories by name
func (r *CategoryRepository) SearchByName(s string) ([]Category, error) {
	var categories []Category
	err := r.db.Where("name LIKE ?", "%"+s+"%").Find(&categories).Error
	return categories, err
}

//Get all categories with products //TODO: Check this
func (r *CategoryRepository) FindAllWithProducts() ([]Category, error) {
	var categories []Category
	err := r.db.Preload("Products").Where("deleted_at = ?", nil).Find(&categories).Error
	return categories, err
}
