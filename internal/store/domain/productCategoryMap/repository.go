package productcategorymap

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type productCategoryMapRepository struct {
	db *gorm.DB
}

//interface
type IProductCategoryMapRepository interface {
	Migration()
	Create(productcategorymap *ProductCategoryMap) error
	Update(productcategorymap *ProductCategoryMap) error
	Delete(productcategorymap *ProductCategoryMap) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]ProductCategoryMap, error)
	FindByID(id uuid.UUID) (*ProductCategoryMap, error)
	FindByProductID(productid uuid.UUID) ([]ProductCategoryMap, error)
	FindByCategoryID(categoryid uuid.UUID) ([]ProductCategoryMap, error)
}

var ProductCategoryMapRepository IProductCategoryMapRepository = &productCategoryMapRepository{}

func NewProductCategoryMapRepository(db *gorm.DB) *productCategoryMapRepository {
	return &productCategoryMapRepository{db: db}
}
func (r *productCategoryMapRepository) Migration() {
	r.db.AutoMigrate(&ProductCategoryMap{})
}

//Create a new productcategorymap
func (r *productCategoryMapRepository) Create(productcategorymap *ProductCategoryMap) error {
	return r.db.Create(productcategorymap).Error
}

//Update a productcategorymap
func (r *productCategoryMapRepository) Update(productcategorymap *ProductCategoryMap) error {
	return r.db.Save(productcategorymap).Error
}

//Delete a productcategorymap
func (r *productCategoryMapRepository) Delete(productcategorymap *ProductCategoryMap) error {
	return r.db.Delete(productcategorymap).Error
}

//Delete a productcategorymap by id
func (r *productCategoryMapRepository) DeleteByID(id uuid.UUID) error {
	productcategorymap := ProductCategoryMap{}
	productcategorymap.ID = id
	return r.db.Delete(&productcategorymap).Error
}

//Find all productcategorymaps
func (r *productCategoryMapRepository) FindAll() ([]ProductCategoryMap, error) {
	var productcategorymaps []ProductCategoryMap
	err := r.db.Find(&productcategorymaps).Error
	return productcategorymaps, err
}

//Find a productcategorymap by id
func (r *productCategoryMapRepository) FindByID(id uuid.UUID) (*ProductCategoryMap, error) {
	productcategorymap := ProductCategoryMap{}
	err := r.db.First(&productcategorymap, id).Error
	return &productcategorymap, err
}

//Find a productcategorymap by productid
func (r *productCategoryMapRepository) FindByProductID(productid uuid.UUID) ([]ProductCategoryMap, error) {
	var productcategorymaps []ProductCategoryMap
	err := r.db.Where("productid = ? AND deleted_at is null", productid).Find(&productcategorymaps).Error
	return productcategorymaps, err
}

//Find a productcategorymap by categoryid
func (r *productCategoryMapRepository) FindByCategoryID(categoryid uuid.UUID) ([]ProductCategoryMap, error) {
	var productcategorymaps []ProductCategoryMap
	err := r.db.Where("categoryid = ? AND deleted_at is null", categoryid).Find(&productcategorymaps).Error
	return productcategorymaps, err
}
