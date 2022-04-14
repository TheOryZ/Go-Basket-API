package services

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/productcategorymap"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/gofrs/uuid"
)

//ProductCategoryMapService is an interface for ProductCategoryMapService
type ProductCategoryMapService interface {
	Create(model dtos.ProductCategoryMapCreateDTO) (dtos.ProductCategoryMapListDTO, error)
	Update(model dtos.ProductCategoryMapUpdateDTO) (dtos.ProductCategoryMapListDTO, error)
	Delete(model dtos.ProductCategoryMapUpdateDTO) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]dtos.ProductCategoryMapListDTO, error)
	FindByID(id uuid.UUID) (dtos.ProductCategoryMapListDTO, error)
	FindByProductID(productID uuid.UUID) ([]dtos.ProductCategoryMapListDTO, error)
	FindByCategoryID(categoryID uuid.UUID) ([]dtos.ProductCategoryMapListDTO, error)
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
func (r *productCategoryMapService) Create(model dtos.ProductCategoryMapCreateDTO) (dtos.ProductCategoryMapListDTO, error) {
	productCategoryMap := dtos.ProductCategoryMapListDTO{}
	productCategoryEntity := productcategorymap.ProductCategoryMap{}
	productCategoryEntity.ID = uuid.Must(uuid.NewV4())
	productCategoryEntity.ProductID = model.ProductID
	productCategoryEntity.CategoryID = model.CategoryID
	err := r.productCategoryMapRepository.Create(&productCategoryEntity)
	if err != nil {
		return productCategoryMap, err
	}
	productCategoryMap.ID = productCategoryEntity.ID
	productCategoryMap.Product.ID = productCategoryEntity.ProductID
	productCategoryMap.Category.ID = productCategoryEntity.CategoryID
	return productCategoryMap, nil
}

//Update a productcategorymap
func (r *productCategoryMapService) Update(model dtos.ProductCategoryMapUpdateDTO) (dtos.ProductCategoryMapListDTO, error) {
	productCategoryMap := dtos.ProductCategoryMapListDTO{}
	productCategoryEntity := productcategorymap.ProductCategoryMap{}
	productCategoryEntity.ID = model.ID
	productCategoryEntity.ProductID = model.ProductID
	productCategoryEntity.CategoryID = model.CategoryID
	err := r.productCategoryMapRepository.Update(&productCategoryEntity)
	if err != nil {
		return productCategoryMap, err
	}
	productCategoryMap.ID = productCategoryEntity.ID
	productCategoryMap.Product.ID = productCategoryEntity.ProductID
	productCategoryMap.Category.ID = productCategoryEntity.CategoryID
	return productCategoryMap, nil
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
func (r *productCategoryMapService) FindAll() ([]dtos.ProductCategoryMapListDTO, error) {
	productCategoryMaps := []dtos.ProductCategoryMapListDTO{}
	productCategoryEntities, err := r.productCategoryMapRepository.FindAll()
	if err != nil {
		return productCategoryMaps, err
	}
	for _, productCategoryEntity := range productCategoryEntities {
		productCategoryMap := dtos.ProductCategoryMapListDTO{}
		productCategoryMap.ID = productCategoryEntity.ID
		productCategoryMap.Product.ID = productCategoryEntity.ProductID
		productCategoryMap.Category.ID = productCategoryEntity.CategoryID
		productCategoryMaps = append(productCategoryMaps, productCategoryMap)
	}
	return productCategoryMaps, nil
}

//FindByID productcategorymap
func (r *productCategoryMapService) FindByID(id uuid.UUID) (dtos.ProductCategoryMapListDTO, error) {
	productCategoryMap := dtos.ProductCategoryMapListDTO{}
	productCategoryEntity, err := r.productCategoryMapRepository.FindByID(id)
	if err != nil {
		return productCategoryMap, err
	}
	productCategoryMap.ID = productCategoryEntity.ID
	productCategoryMap.Product.ID = productCategoryEntity.ProductID
	productCategoryMap.Category.ID = productCategoryEntity.CategoryID
	return productCategoryMap, nil
}

//FindByProductID productcategorymap
func (r *productCategoryMapService) FindByProductID(productID uuid.UUID) ([]dtos.ProductCategoryMapListDTO, error) {
	productCategoryMaps := []dtos.ProductCategoryMapListDTO{}
	productCategoryEntities, err := r.productCategoryMapRepository.FindByProductID(productID)
	if err != nil {
		return productCategoryMaps, err
	}
	for _, productCategoryEntity := range productCategoryEntities {
		productCategoryMap := dtos.ProductCategoryMapListDTO{}
		productCategoryMap.ID = productCategoryEntity.ID
		productCategoryMap.Product.ID = productCategoryEntity.ProductID
		productCategoryMap.Category.ID = productCategoryEntity.CategoryID
		productCategoryMaps = append(productCategoryMaps, productCategoryMap)
	}
	return productCategoryMaps, nil
}

//FindByCategoryID productcategorymap
func (r *productCategoryMapService) FindByCategoryID(categoryID uuid.UUID) ([]dtos.ProductCategoryMapListDTO, error) {
	productCategoryMaps := []dtos.ProductCategoryMapListDTO{}
	productCategoryEntities, err := r.productCategoryMapRepository.FindByCategoryID(categoryID)
	if err != nil {
		return productCategoryMaps, err
	}
	for _, productCategoryEntity := range productCategoryEntities {
		productCategoryMap := dtos.ProductCategoryMapListDTO{}
		productCategoryMap.ID = productCategoryEntity.ID
		productCategoryMap.Product.ID = productCategoryEntity.ProductID
		productCategoryMap.Category.ID = productCategoryEntity.CategoryID
		productCategoryMaps = append(productCategoryMaps, productCategoryMap)
	}
	return productCategoryMaps, nil
}
