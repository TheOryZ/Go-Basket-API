package services

import (
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/status"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/gofrs/uuid"
)

// StatusService is an interface for StatusService
type StatusService interface {
	Create(model dtos.StatusCreateDTO) (dtos.StatusListDTO, error)
	Update(model dtos.StatusUpdateDTO) (dtos.StatusListDTO, error)
	Delete(model dtos.StatusUpdateDTO) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]dtos.StatusListDTO, error)
	FindByID(id uuid.UUID) (dtos.StatusListDTO, error)
	FindByName(name string) (dtos.StatusListDTO, error)
}

//statusService is an implementation of StatusService
type statusService struct {
	statusRepository status.IStatusRepository
}

//NewStatusService is a constructor for StatusService
func NewStatusService(statusRepository status.IStatusRepository) StatusService {
	return &statusService{statusRepository: statusRepository}
}

//Create a new status
func (s *statusService) Create(model dtos.StatusCreateDTO) (dtos.StatusListDTO, error) {
	listModel := dtos.StatusListDTO{}
	statusModel := status.Status{
		Name:      model.Name,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := s.statusRepository.Create(&statusModel)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.StatusListDTO{
		ID:   statusModel.ID,
		Name: statusModel.Name,
	}
	return listModel, nil
}

//Update a status
func (s *statusService) Update(model dtos.StatusUpdateDTO) (dtos.StatusListDTO, error) {
	listModel := dtos.StatusListDTO{}
	statusModel := status.Status{
		ID:        model.ID,
		Name:      model.Name,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := s.statusRepository.Update(&statusModel)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.StatusListDTO{
		ID:   statusModel.ID,
		Name: statusModel.Name,
	}
	return listModel, nil
}

//Delete a status
func (s *statusService) Delete(model dtos.StatusUpdateDTO) error {
	statusModel := status.Status{
		ID: model.ID,
	}
	err := s.statusRepository.Delete(&statusModel)
	return err
}

//Delete a status by id
func (s *statusService) DeleteByID(id uuid.UUID) error {
	err := s.statusRepository.DeleteByID(id)
	return err
}

//FindAll statuses
func (s *statusService) FindAll() ([]dtos.StatusListDTO, error) {
	listModel := []dtos.StatusListDTO{}
	statusModels, err := s.statusRepository.FindAll()
	if err != nil {
		return listModel, err
	}
	for _, statusModel := range statusModels {
		listModel = append(listModel, dtos.StatusListDTO{
			ID:   statusModel.ID,
			Name: statusModel.Name,
		})
	}
	return listModel, nil
}

//FindByID a status by id
func (s *statusService) FindByID(id uuid.UUID) (dtos.StatusListDTO, error) {
	listModel := dtos.StatusListDTO{}
	statusModel, err := s.statusRepository.FindByID(id)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.StatusListDTO{
		ID:   statusModel.ID,
		Name: statusModel.Name,
	}
	return listModel, nil
}

//FindByName a status by name
func (s *statusService) FindByName(name string) (dtos.StatusListDTO, error) {
	listModel := dtos.StatusListDTO{}
	statusModel, err := s.statusRepository.FindByName(name)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.StatusListDTO{
		ID:   statusModel.ID,
		Name: statusModel.Name,
	}
	return listModel, nil
}
