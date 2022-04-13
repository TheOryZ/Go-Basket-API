package services

import (
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/category"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/gofrs/uuid"
)

//CategoryService is an interface for CategoryService
type CategoryService interface {
	Create(model dtos.CategoryCreateDTO) (dtos.CategoryListDTO, error)
	Update(model dtos.CategoryUpdateDTO) (dtos.CategoryListDTO, error)
	Delete(model dtos.CategoryUpdateDTO) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]dtos.CategoryListDTO, error)
	FindAllWithPagination(page int, limit int) ([]dtos.CategoryListDTO, error)
	CountAll() (int64, error)
	FindByID(id uuid.UUID) (dtos.CategoryListDTO, error)
	FindByName(name string) (dtos.CategoryListDTO, error)
	SearchByName(ss string) ([]dtos.CategoryListDTO, error)
	FindAllWithProducts() ([]dtos.CategoryListDTO, error) //TODO:
}

//categoryService is an implementation of CategoryService
type categoryService struct {
	categoryRepository category.ICategoryRepository
}

//NewCategoryService is a constructor for CategoryService
func NewCategoryService(categoryRepository category.ICategoryRepository) CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}

//Create a new category
func (s *categoryService) Create(model dtos.CategoryCreateDTO) (dtos.CategoryListDTO, error) {
	listModel := dtos.CategoryListDTO{}
	categoryModel := category.Category{
		Name:      model.Name,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		IsActive:  true,
	}
	err := s.categoryRepository.Create(&categoryModel)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.CategoryListDTO{
		ID:   categoryModel.ID,
		Name: categoryModel.Name,
	}
	return listModel, nil
}

//Update a category
func (s *categoryService) Update(model dtos.CategoryUpdateDTO) (dtos.CategoryListDTO, error) {
	listModel := dtos.CategoryListDTO{}
	categoryModel := category.Category{
		ID:        model.ID,
		Name:      model.Name,
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := s.categoryRepository.Update(&categoryModel)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.CategoryListDTO{
		ID:   categoryModel.ID,
		Name: categoryModel.Name,
	}
	return listModel, nil
}

//Delete a category
func (s *categoryService) Delete(model dtos.CategoryUpdateDTO) error {
	categoryModel, err := s.categoryRepository.FindByID(model.ID)
	if err != nil {
		return err
	}
	err = s.categoryRepository.Delete(categoryModel)
	if err != nil {
		return err
	}
	return nil
}

//DeleteByID a category
func (s *categoryService) DeleteByID(id uuid.UUID) error {
	categoryModel, err := s.categoryRepository.FindByID(id)
	if err != nil {
		return err
	}
	err = s.categoryRepository.Delete(categoryModel)
	if err != nil {
		return err
	}
	return nil
}

//FindAll category
func (s *categoryService) FindAll() ([]dtos.CategoryListDTO, error) {
	listModel := []dtos.CategoryListDTO{}
	categoryModels, err := s.categoryRepository.FindAll()
	if err != nil {
		return listModel, err
	}
	for _, categoryModel := range categoryModels {
		listModel = append(listModel, dtos.CategoryListDTO{
			ID:   categoryModel.ID,
			Name: categoryModel.Name,
		})
	}
	return listModel, nil
}

//FindAllWithPagination category
func (s *categoryService) FindAllWithPagination(page int, limit int) ([]dtos.CategoryListDTO, error) {
	listModel := []dtos.CategoryListDTO{}
	categoryModels, err := s.categoryRepository.FindAllWithPagination(page, limit)
	if err != nil {
		return listModel, err
	}
	for _, categoryModel := range categoryModels {
		listModel = append(listModel, dtos.CategoryListDTO{
			ID:   categoryModel.ID,
			Name: categoryModel.Name,
		})
	}
	return listModel, nil
}

//CountAll category
func (s *categoryService) CountAll() (int64, error) {
	count, err := s.categoryRepository.CountAll()
	if err != nil {
		return 0, err
	}
	return count, nil
}

//FindByID category
func (s *categoryService) FindByID(id uuid.UUID) (dtos.CategoryListDTO, error) {
	listModel := dtos.CategoryListDTO{}
	categoryModel, err := s.categoryRepository.FindByID(id)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.CategoryListDTO{
		ID:   categoryModel.ID,
		Name: categoryModel.Name,
	}
	return listModel, nil
}

//FindByName category
func (s *categoryService) FindByName(name string) (dtos.CategoryListDTO, error) {
	listModel := dtos.CategoryListDTO{}
	categoryModel, err := s.categoryRepository.FindByName(name)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.CategoryListDTO{
		ID:   categoryModel.ID,
		Name: categoryModel.Name,
	}
	return listModel, nil
}

//SearchByName category
func (s *categoryService) SearchByName(ss string) ([]dtos.CategoryListDTO, error) {
	listModel := []dtos.CategoryListDTO{}
	categoryModels, err := s.categoryRepository.SearchByName(ss)
	if err != nil {
		return listModel, err
	}
	for _, categoryModel := range categoryModels {
		listModel = append(listModel, dtos.CategoryListDTO{
			ID:   categoryModel.ID,
			Name: categoryModel.Name,
		})
	}
	return listModel, nil
}

//FindAllWithProducts category
func (s *categoryService) FindAllWithProducts() ([]dtos.CategoryListDTO, error) {
	listModel := []dtos.CategoryListDTO{}
	categoryModels, err := s.categoryRepository.FindAllWithProducts()
	if err != nil {
		return listModel, err
	}
	for _, categoryModel := range categoryModels {
		listModel = append(listModel, dtos.CategoryListDTO{
			ID:   categoryModel.ID,
			Name: categoryModel.Name,
		})
	}
	return listModel, nil
}
