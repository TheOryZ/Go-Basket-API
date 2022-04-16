package services

import (
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/product"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/gofrs/uuid"
)

// ProductService is an interface for ProductService
type ProductService interface {
	Create(model dtos.ProductCreateDTO) (dtos.ProductListDTO, error)
	Update(model dtos.ProductUpdateDTO) (dtos.ProductListDTO, error)
	Delete(model dtos.ProductUpdateDTO) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]dtos.ProductListDTO, error)
	FindAllWithPagination(page int, limit int) ([]dtos.ProductListDTO, error)
	CountAll() (int64, error)
	FindByID(id uuid.UUID) (dtos.ProductListDTO, error)
	FindByCategoryID(id uuid.UUID) ([]dtos.ProductListDTO, error)
	SearchByName(name string) ([]dtos.ProductListDTO, error)
}

//productService is an implementation of ProductService
type productService struct {
	productRepository product.IProductRepository
}

//NewProductService is a constructor for ProductService
func NewProductService(productRepository product.IProductRepository) ProductService {
	return &productService{productRepository: productRepository}
}

//Create a new product
func (s *productService) Create(model dtos.ProductCreateDTO) (dtos.ProductListDTO, error) {
	listModel := dtos.ProductListDTO{}
	productModel := product.Product{
		Name:             model.Name,
		SKU:              model.SKU,
		ShortDescription: model.ShortDescription,
		Description:      model.Description,
		Price:            model.Price,
		UnitOfStock:      model.UnitOfStock,
		CreatedAt:        time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:        time.Now().Format("2006-01-02 15:04:05"),
		IsActive:         true,
	}
	err := s.productRepository.Create(&productModel)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.ProductListDTO{
		ID:               productModel.ID,
		Name:             productModel.Name,
		SKU:              productModel.SKU,
		ShortDescription: productModel.ShortDescription,
		Description:      productModel.Description,
		Price:            productModel.Price,
		UnitOfStock:      productModel.UnitOfStock,
	}
	return listModel, nil
}

//Update a product
func (s *productService) Update(model dtos.ProductUpdateDTO) (dtos.ProductListDTO, error) {
	listModel := dtos.ProductListDTO{}
	productModel := product.Product{
		ID:               model.ID,
		Name:             model.Name,
		SKU:              model.SKU,
		ShortDescription: model.ShortDescription,
		Description:      model.Description,
		Price:            model.Price,
		UnitOfStock:      model.UnitOfStock,
		CreatedAt:        time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:        time.Now().Format("2006-01-02 15:04:05"),
		IsActive:         true,
	}
	err := s.productRepository.Update(&productModel)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.ProductListDTO{
		ID:   productModel.ID,
		Name: productModel.Name,
	}
	return listModel, nil
}

//Delete a product
func (s *productService) Delete(model dtos.ProductUpdateDTO) error {
	productModel := product.Product{
		ID:               model.ID,
		Name:             model.Name,
		SKU:              model.SKU,
		ShortDescription: model.ShortDescription,
		Description:      model.Description,
		Price:            model.Price,
		UnitOfStock:      model.UnitOfStock,
		CreatedAt:        time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:        time.Now().Format("2006-01-02 15:04:05"),
		IsActive:         true,
	}
	err := s.productRepository.Delete(&productModel)
	return err
}

//DeleteByID a product
func (s *productService) DeleteByID(id uuid.UUID) error {
	productModel := product.Product{
		ID: id,
	}
	err := s.productRepository.Delete(&productModel)
	return err
}

//FindAll products
func (s *productService) FindAll() ([]dtos.ProductListDTO, error) {
	listModel := []dtos.ProductListDTO{}
	products, err := s.productRepository.FindAll()
	if err != nil {
		return listModel, err
	}
	for _, product := range products {
		listModel = append(listModel, dtos.ProductListDTO{
			ID:               product.ID,
			Name:             product.Name,
			SKU:              product.SKU,
			ShortDescription: product.ShortDescription,
			Description:      product.Description,
			Price:            product.Price,
			UnitOfStock:      product.UnitOfStock,
		})
	}
	return listModel, nil
}

//FindAllWithPagination products
func (s *productService) FindAllWithPagination(page int, limit int) ([]dtos.ProductListDTO, error) {
	listModel := []dtos.ProductListDTO{}
	products, err := s.productRepository.FindAllWithPagination(page, limit)
	if err != nil {
		return listModel, err
	}
	for _, product := range products {
		listModel = append(listModel, dtos.ProductListDTO{
			ID:               product.ID,
			Name:             product.Name,
			SKU:              product.SKU,
			ShortDescription: product.ShortDescription,
			Description:      product.Description,
			Price:            product.Price,
			UnitOfStock:      product.UnitOfStock,
		})
	}
	return listModel, nil
}

//CountAll products
func (s *productService) CountAll() (int64, error) {
	count, err := s.productRepository.CountAll()
	return count, err
}

//FindByID a product
func (s *productService) FindByID(id uuid.UUID) (dtos.ProductListDTO, error) {
	listModel := dtos.ProductListDTO{}
	productModel, err := s.productRepository.FindByID(id)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.ProductListDTO{
		ID:               productModel.ID,
		Name:             productModel.Name,
		SKU:              productModel.SKU,
		ShortDescription: productModel.ShortDescription,
		Description:      productModel.Description,
		Price:            productModel.Price,
		UnitOfStock:      productModel.UnitOfStock,
	}
	return listModel, nil
}

//FindByCategoryID a product
func (s *productService) FindByCategoryID(id uuid.UUID) ([]dtos.ProductListDTO, error) {
	listModel := []dtos.ProductListDTO{}
	products, err := s.productRepository.FindByCategoryID(id)
	if err != nil {
		return listModel, err
	}
	for _, product := range products {
		listModel = append(listModel, dtos.ProductListDTO{
			ID:               product.ID,
			Name:             product.Name,
			SKU:              product.SKU,
			ShortDescription: product.ShortDescription,
			Description:      product.Description,
			Price:            product.Price,
			UnitOfStock:      product.UnitOfStock,
		})
	}
	return listModel, nil
}

//FindByName a product
func (s *productService) SearchByName(name string) ([]dtos.ProductListDTO, error) {
	listModel := []dtos.ProductListDTO{}
	products, err := s.productRepository.SearchByName(name)
	if err != nil {
		return listModel, err
	}
	for _, product := range products {
		listModel = append(listModel, dtos.ProductListDTO{
			ID:               product.ID,
			Name:             product.Name,
			SKU:              product.SKU,
			ShortDescription: product.ShortDescription,
			Description:      product.Description,
			Price:            product.Price,
			UnitOfStock:      product.UnitOfStock,
		})
	}
	return listModel, nil
}
