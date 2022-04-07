package product

import "gorm.io/gorm"

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
func (r *ProductRepository) Migration() {
	r.db.AutoMigrate(&Product{})
}

//Create a new product
func (r *ProductRepository) Create(product *Product) error {
	return r.db.Create(product).Error
}

//Update a product
func (r *ProductRepository) Update(product *Product) error {
	return r.db.Save(product).Error
}

//Delete a product
func (r *ProductRepository) Delete(product *Product) error {
	return r.db.Delete(product).Error
}

//Delete a product by id
func (r *ProductRepository) DeleteByID(id uint) error {
	product := Product{}
	product.ID = id
	return r.db.Delete(&product).Error
}

//Find all products
func (r *ProductRepository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Where("deleted_at = ?", nil).Find(&products).Error
	return products, err
}

//Find all with pagination  //TODO: check this
func (r *ProductRepository) FindAllWithPagination(page int, limit int) ([]Product, error) {
	var products []Product
	err := r.db.Where("deleted_at = ?", nil).Offset(page).Limit(limit).Find(&products).Error
	return products, err
}

//Count all products
func (r *ProductRepository) CountAll() (int64, error) {
	var count int64
	err := r.db.Model(&Product{}).Where("deleted_at = ?", nil).Count(&count).Error
	return count, err
}

//Find a product by id
func (r *ProductRepository) FindByID(id uint) (*Product, error) {
	product := Product{}
	err := r.db.First(&product, id).Error
	return &product, err
}

//Search products by name or SKU
func (r *ProductRepository) SearchByName(s string) ([]Product, error) {
	var products []Product
	err := r.db.Where("name LIKE ? OR SKU LIKE ?", "%"+s+"%", "%"+s+"%").Find(&products).Error
	return products, err
}
