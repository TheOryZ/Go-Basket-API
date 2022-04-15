package product

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

//interface
type IProductRepository interface {
	Migration()
	Create(product *Product) error
	Update(product *Product) error
	Delete(product *Product) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]Product, error)
	FindAllWithPagination(page int, limit int) ([]Product, error)
	CountAll() (int64, error)
	FindByID(id uuid.UUID) (*Product, error)
	FindByCategoryID(id uuid.UUID) ([]Product, error)
	SearchByName(s string) ([]Product, error)
}

var ProductRepository IProductRepository = &productRepository{}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}
func (r *productRepository) Migration() {
	r.db.AutoMigrate(&Product{})
}

//Create a new product
func (r *productRepository) Create(product *Product) error {
	return r.db.Create(product).Error
}

//Update a product
func (r *productRepository) Update(product *Product) error {
	return r.db.Save(product).Error
}

//Delete a product
func (r *productRepository) Delete(product *Product) error {
	return r.db.Delete(product).Error
}

//Delete a product by id
func (r *productRepository) DeleteByID(id uuid.UUID) error {
	product := Product{}
	product.ID = id
	return r.db.Delete(&product).Error
}

//Find all products
func (r *productRepository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Where("deleted_at = ?", nil).Find(&products).Error
	return products, err
}

//Find all with pagination  //TODO: check this
func (r *productRepository) FindAllWithPagination(page int, limit int) ([]Product, error) {
	var products []Product
	err := r.db.Where("deleted_at = ?", nil).Offset(page).Limit(limit).Find(&products).Error
	return products, err
}

//Count all products
func (r *productRepository) CountAll() (int64, error) {
	var count int64
	err := r.db.Model(&Product{}).Where("deleted_at = ?", nil).Count(&count).Error
	return count, err
}

//Find a product by id
func (r *productRepository) FindByID(id uuid.UUID) (*Product, error) {
	product := Product{}
	err := r.db.First(&product, id).Error
	return &product, err
}

//FindByCategoryID find products by category id
func (r *productRepository) FindByCategoryID(id uuid.UUID) ([]Product, error) {
	var products []Product
	err := r.db.Joins("INNER JOIN product_category_map map on map.product_id = products.id").
		Where("map.category_id = ?", id).
		Where("deleted_at = ?", nil).
		Select("products.*").
		Find(&products).Error
	return products, err
}

//Search products by name or SKU
func (r *productRepository) SearchByName(s string) ([]Product, error) {
	var products []Product
	err := r.db.Where("name LIKE ? OR SKU LIKE ?", "%"+s+"%", "%"+s+"%").Find(&products).Error
	return products, err
}
