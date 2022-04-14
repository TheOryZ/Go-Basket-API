package services

import (
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/userrolemap"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/gofrs/uuid"
)

// UserRoleMapService is an interface for UserRoleMapService
type UserRoleMapService interface {
	Create(model dtos.UserRoleMapCreateDTO) (dtos.UserRoleMapListDTO, error)
	Update(model dtos.UserRoleMapUpdateDTO) (dtos.UserRoleMapListDTO, error)
	Delete(model dtos.UserRoleMapUpdateDTO) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]dtos.UserRoleMapListDTO, error)
	FindByID(id uuid.UUID) (dtos.UserRoleMapListDTO, error)
	FindByUserID(userID uuid.UUID) ([]dtos.UserRoleMapListDTO, error)
	FindByRoleID(roleID uuid.UUID) ([]dtos.UserRoleMapListDTO, error)
}

// userRoleMapService is a struct for UserRoleMapService
type userRoleMapService struct {
	userRoleMapRepository userrolemap.IUserRoleMapRepository
}

// NewUserRoleMapService is a constructor for UserRoleMapService
func NewUserRoleMapService(userRoleMapRepository userrolemap.IUserRoleMapRepository) UserRoleMapService {
	return &userRoleMapService{userRoleMapRepository: userRoleMapRepository}
}

// Create a new userrolemap
func (r *userRoleMapService) Create(model dtos.UserRoleMapCreateDTO) (dtos.UserRoleMapListDTO, error) {
	userRoleMap := dtos.UserRoleMapListDTO{}
	userRoleEntity := userrolemap.UserRoleMap{}
	userRoleEntity.ID = uuid.Must(uuid.NewV4())
	userRoleEntity.UserID = model.UserID
	userRoleEntity.RoleID = model.RoleID
	userRoleEntity.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	userRoleEntity.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	userRoleEntity.IsActive = true
	err := r.userRoleMapRepository.Create(&userRoleEntity)
	if err != nil {
		return userRoleMap, err
	}
	userRoleMap.ID = userRoleEntity.ID
	userRoleMap.User.ID = userRoleEntity.UserID
	userRoleMap.Role.ID = userRoleEntity.RoleID
	return userRoleMap, nil
}

// Update a userrolemap
func (r *userRoleMapService) Update(model dtos.UserRoleMapUpdateDTO) (dtos.UserRoleMapListDTO, error) {
	userRoleMap := dtos.UserRoleMapListDTO{}
	userRoleEntity := userrolemap.UserRoleMap{}
	userRoleEntity.ID = model.ID
	userRoleEntity.UserID = model.UserID
	userRoleEntity.RoleID = model.RoleID
	err := r.userRoleMapRepository.Update(&userRoleEntity)
	if err != nil {
		return userRoleMap, err
	}
	userRoleMap.ID = userRoleEntity.ID
	userRoleMap.User.ID = userRoleEntity.UserID
	userRoleMap.Role.ID = userRoleEntity.RoleID
	return userRoleMap, nil
}

// Delete a userrolemap
func (r *userRoleMapService) Delete(model dtos.UserRoleMapUpdateDTO) error {
	userRoleEntity := userrolemap.UserRoleMap{}
	userRoleEntity.ID = model.ID
	err := r.userRoleMapRepository.Delete(&userRoleEntity)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByID a userrolemap
func (r *userRoleMapService) DeleteByID(id uuid.UUID) error {
	userRoleEntity := userrolemap.UserRoleMap{}
	userRoleEntity.ID = id
	err := r.userRoleMapRepository.Delete(&userRoleEntity)
	if err != nil {
		return err
	}
	return nil
}

// FindAll userrolemaps
func (r *userRoleMapService) FindAll() ([]dtos.UserRoleMapListDTO, error) {
	userRoleMapList := []dtos.UserRoleMapListDTO{}
	userRoleMapEntities, err := r.userRoleMapRepository.FindAll()
	if err != nil {
		return userRoleMapList, err
	}
	for _, userRoleMapEntity := range userRoleMapEntities {
		userRoleMap := dtos.UserRoleMapListDTO{}
		userRoleMap.ID = userRoleMapEntity.ID
		userRoleMap.User.ID = userRoleMapEntity.UserID
		userRoleMap.Role.ID = userRoleMapEntity.RoleID
		userRoleMapList = append(userRoleMapList, userRoleMap)
	}
	return userRoleMapList, nil
}

// FindByID userrolemaps
func (r *userRoleMapService) FindByID(id uuid.UUID) (dtos.UserRoleMapListDTO, error) {
	userRoleMap := dtos.UserRoleMapListDTO{}
	userRoleMapEntity, err := r.userRoleMapRepository.FindByID(id)
	if err != nil {
		return userRoleMap, err
	}
	userRoleMap.ID = userRoleMapEntity.ID
	userRoleMap.User.ID = userRoleMapEntity.UserID
	userRoleMap.Role.ID = userRoleMapEntity.RoleID
	return userRoleMap, nil
}

// FindByUserID userrolemaps
func (r *userRoleMapService) FindByUserID(userID uuid.UUID) ([]dtos.UserRoleMapListDTO, error) {
	userRoleMapList := []dtos.UserRoleMapListDTO{}
	userRoleMapEntities, err := r.userRoleMapRepository.FindByUserID(userID)
	if err != nil {
		return userRoleMapList, err
	}
	for _, userRoleMapEntity := range *userRoleMapEntities {
		userRoleMap := dtos.UserRoleMapListDTO{}
		userRoleMap.ID = userRoleMapEntity.ID
		userRoleMap.User.ID = userRoleMapEntity.UserID
		userRoleMap.Role.ID = userRoleMapEntity.RoleID
		userRoleMapList = append(userRoleMapList, userRoleMap)
	}
	return userRoleMapList, nil
}

// FindByRoleID userrolemaps
func (r *userRoleMapService) FindByRoleID(roleID uuid.UUID) ([]dtos.UserRoleMapListDTO, error) {
	userRoleMapList := []dtos.UserRoleMapListDTO{}
	userRoleMapEntities, err := r.userRoleMapRepository.FindByRoleID(roleID)
	if err != nil {
		return userRoleMapList, err
	}
	for _, userRoleMapEntity := range *userRoleMapEntities {
		userRoleMap := dtos.UserRoleMapListDTO{}
		userRoleMap.ID = userRoleMapEntity.ID
		userRoleMap.User.ID = userRoleMapEntity.UserID
		userRoleMap.Role.ID = userRoleMapEntity.RoleID
		userRoleMapList = append(userRoleMapList, userRoleMap)
	}
	return userRoleMapList, nil
}
