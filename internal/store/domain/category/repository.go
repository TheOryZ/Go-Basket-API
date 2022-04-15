package category

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

//interface
type ICategoryRepository interface {
	Migration()
	Create(category *Category) error
	Update(category *Category) error
	Delete(category *Category) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]Category, error)
	FindAllWithPagination(page int, limit int) ([]Category, error)
	CountAll() (int64, error)
	FindByID(id uuid.UUID) (*Category, error)
	FindByName(name string) (*Category, error)
	SearchByName(s string) ([]Category, error)
	FindByProductID(id uuid.UUID) ([]Category, error)
	FindAllWithProducts() ([]Category, error)
}

var CategoryRepository ICategoryRepository = &categoryRepository{}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}
func (r *categoryRepository) Migration() {
	r.db.AutoMigrate(&Category{})
}

//Create a new category
func (r *categoryRepository) Create(category *Category) error {
	return r.db.Create(category).Error
}

//Update a category
func (r *categoryRepository) Update(category *Category) error {
	return r.db.Save(category).Error
}

//Delete a category
func (r *categoryRepository) Delete(category *Category) error {
	return r.db.Delete(category).Error
}

//Delete a category by id
func (r *categoryRepository) DeleteByID(id uuid.UUID) error {
	category := Category{}
	category.ID = id
	return r.db.Delete(&category).Error
}

//Find all categories
func (r *categoryRepository) FindAll() ([]Category, error) {
	var categories []Category
	err := r.db.Where("deleted_at = ?", nil).Find(&categories).Error
	return categories, err
}

//Find all with pagination  //TODO: check this
func (r *categoryRepository) FindAllWithPagination(page int, limit int) ([]Category, error) {
	var categories []Category
	err := r.db.Where("deleted_at = ?", nil).Offset(page).Limit(limit).Find(&categories).Error
	return categories, err
}

//Count all categories
func (r *categoryRepository) CountAll() (int64, error) {
	var count int64
	err := r.db.Model(&Category{}).Where("deleted_at = ?", nil).Count(&count).Error
	return count, err
}

//Find a category by id
func (r *categoryRepository) FindByID(id uuid.UUID) (*Category, error) {
	category := Category{}
	err := r.db.First(&category, id).Error
	return &category, err
}

//Find a category by name
func (r *categoryRepository) FindByName(name string) (*Category, error) {
	category := Category{}
	err := r.db.Where("name = ?", name).First(&category).Error
	return &category, err
}

//Search categories by name
func (r *categoryRepository) SearchByName(s string) ([]Category, error) {
	var categories []Category
	err := r.db.Where("name LIKE ?", "%"+s+"%").Find(&categories).Error
	return categories, err
}

//FindByProductID find all categories by product id
func (r *categoryRepository) FindByProductID(id uuid.UUID) ([]Category, error) {
	var categories []Category
	err := r.db.Joins(" INNER JOIN product_category_map map ON map.category_id = category.id").
		Where("map.product_id = ?", id).
		Where("category.deleted_at = ?", nil).
		Select("category.*").
		Find(&categories).Error
	return categories, err
}

//Get all categories with products //TODO: Check this
func (r *categoryRepository) FindAllWithProducts() ([]Category, error) {
	var categories []Category
	err := r.db.Joins(" INNER JOIN product_category_map map ON map.category_id = category.id").
		Where("category.deleted_at = ?", nil).
		Select("category.*").
		Find(&categories).Error
	return categories, err
}
