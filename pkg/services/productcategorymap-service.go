package services

import (
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/productcategorymap"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/gofrs/uuid"
)

//ProductCategoryMapService is an interface for ProductCategoryMapService
type ProductCategoryMapService interface {
	Create(model dtos.ProductCategoryMapCreateDTO) (productcategorymap.ProductCategoryMap, error)
	Update(model dtos.ProductCategoryMapUpdateDTO) (productcategorymap.ProductCategoryMap, error)
	Delete(model dtos.ProductCategoryMapUpdateDTO) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]productcategorymap.ProductCategoryMap, error)
	FindByID(id uuid.UUID) (productcategorymap.ProductCategoryMap, error)
	FindByProductID(productID uuid.UUID) ([]productcategorymap.ProductCategoryMap, error)
	FindByCategoryID(categoryID uuid.UUID) ([]productcategorymap.ProductCategoryMap, error)
}

//productCategoryMapService is a struct for ProductCategoryMapService
type productCategoryMapService struct {
	productCategoryMapRepository productcategorymap.IProductCategoryMapRepository
}

//NewProductCategoryMapService is a constructor for ProductCategoryMapService
func NewProductCategoryMapService(productCategoryMapRepository productcategorymap.IProductCategoryMapRepository) ProductCategoryMapService {
	return &productCategoryMapService{productCategoryMapRepository: productCategoryMapRepository}
}

//Create a new productcategorymap
func (r *productCategoryMapService) Create(model dtos.ProductCategoryMapCreateDTO) (productcategorymap.ProductCategoryMap, error) {
	productCategoryEntity := productcategorymap.ProductCategoryMap{}
	productCategoryEntity.ProductID = model.ProductID
	productCategoryEntity.CategoryID = model.CategoryID
	productCategoryEntity.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	productCategoryEntity.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	productCategoryEntity.IsActive = true
	err := r.productCategoryMapRepository.Create(&productCategoryEntity)
	if err != nil {
		return productCategoryEntity, err
	}
	return productCategoryEntity, nil
}

//Update a productcategorymap
func (r *productCategoryMapService) Update(model dtos.ProductCategoryMapUpdateDTO) (productcategorymap.ProductCategoryMap, error) {
	productCategoryEntity := productcategorymap.ProductCategoryMap{}
	productCategoryEntity.ID = model.ID
	productCategoryEntity.ProductID = model.ProductID
	productCategoryEntity.CategoryID = model.CategoryID
	productCategoryEntity.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	productCategoryEntity.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	productCategoryEntity.IsActive = true
	err := r.productCategoryMapRepository.Update(&productCategoryEntity)
	if err != nil {
		return productCategoryEntity, err
	}
	return productCategoryEntity, nil
}

//Delete a productcategorymap
func (r *productCategoryMapService) Delete(model dtos.ProductCategoryMapUpdateDTO) error {
	productCategoryEntity := productcategorymap.ProductCategoryMap{}
	productCategoryEntity.ID = model.ID
	err := r.productCategoryMapRepository.Delete(&productCategoryEntity)
	if err != nil {
		return err
	}
	return nil
}

//DeleteByID a productcategorymap
func (r *productCategoryMapService) DeleteByID(id uuid.UUID) error {
	productCategoryEntity := productcategorymap.ProductCategoryMap{}
	productCategoryEntity.ID = id
	err := r.productCategoryMapRepository.Delete(&productCategoryEntity)
	if err != nil {
		return err
	}
	return nil
}

//FindAll productcategorymaps
func (r *productCategoryMapService) FindAll() ([]productcategorymap.ProductCategoryMap, error) {
	productcategorymapEmpty := []productcategorymap.ProductCategoryMap{}
	productCategoryEntities, err := r.productCategoryMapRepository.FindAll()
	if err != nil {
		return productcategorymapEmpty, err
	}
	return productCategoryEntities, nil
}

//FindByID productcategorymap
func (r *productCategoryMapService) FindByID(id uuid.UUID) (productcategorymap.ProductCategoryMap, error) {
	productcategorymapEmpty := productcategorymap.ProductCategoryMap{}
	productCategoryEntity, err := r.productCategoryMapRepository.FindByID(id)
	if err != nil {
		return productcategorymapEmpty, err
	}
	return *productCategoryEntity, nil
}

//FindByProductID productcategorymap
func (r *productCategoryMapService) FindByProductID(productID uuid.UUID) ([]productcategorymap.ProductCategoryMap, error) {
	productcategorymapEmpty := []productcategorymap.ProductCategoryMap{}
	productCategoryEntities, err := r.productCategoryMapRepository.FindByProductID(productID)
	if err != nil {
		return productcategorymapEmpty, err
	}
	return productCategoryEntities, nil
}

//FindByCategoryID productcategorymap
func (r *productCategoryMapService) FindByCategoryID(categoryID uuid.UUID) ([]productcategorymap.ProductCategoryMap, error) {
	productcategorymapEmpty := []productcategorymap.ProductCategoryMap{}
	productCategoryEntities, err := r.productCategoryMapRepository.FindByCategoryID(categoryID)
	if err != nil {
		return productcategorymapEmpty, err
	}
	return productCategoryEntities, nil
}
