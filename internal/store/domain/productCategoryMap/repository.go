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

//Create a new productcategorymap
func (r *ProductCategoryMapRepository) Create(productcategorymap *ProductCategoryMap) error {
	return r.db.Create(productcategorymap).Error
}

//Update a productcategorymap
func (r *ProductCategoryMapRepository) Update(productcategorymap *ProductCategoryMap) error {
	return r.db.Save(productcategorymap).Error
}

//Delete a productcategorymap
func (r *ProductCategoryMapRepository) Delete(productcategorymap *ProductCategoryMap) error {
	return r.db.Delete(productcategorymap).Error
}

//Delete a productcategorymap by id
func (r *ProductCategoryMapRepository) DeleteByID(id uint) error {
	productcategorymap := ProductCategoryMap{}
	productcategorymap.ID = id
	return r.db.Delete(&productcategorymap).Error
}

//Find all productcategorymaps
func (r *ProductCategoryMapRepository) FindAll() ([]ProductCategoryMap, error) {
	var productcategorymaps []ProductCategoryMap
	err := r.db.Find(&productcategorymaps).Error
	return productcategorymaps, err
}

//Find a productcategorymap by id
func (r *ProductCategoryMapRepository) FindByID(id uint) (*ProductCategoryMap, error) {
	productcategorymap := ProductCategoryMap{}
	err := r.db.First(&productcategorymap, id).Error
	return &productcategorymap, err
}

//Find a productcategorymap by productid
func (r *ProductCategoryMapRepository) FindByProductID(productid uint) ([]ProductCategoryMap, error) {
	var productcategorymaps []ProductCategoryMap
	err := r.db.Where("productid = ? AND deleted_at = ?", productid, nil).Find(&productcategorymaps).Error
	return productcategorymaps, err
}

//Find a productcategorymap by categoryid
func (r *ProductCategoryMapRepository) FindByCategoryID(categoryid uint) ([]ProductCategoryMap, error) {
	var productcategorymaps []ProductCategoryMap
	err := r.db.Where("categoryid = ? AND deleted_at = ?", categoryid, nil).Find(&productcategorymaps).Error
	return productcategorymaps, err
}
