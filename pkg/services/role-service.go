package services

import (
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/role"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/gofrs/uuid"
)

//RoleService is an interface for RoleService
type RoleService interface {
	Create(model dtos.RoleCreateDTO) (dtos.RoleListDTO, error)
	Update(model dtos.RoleUpdateDTO) (dtos.RoleListDTO, error)
	Delete(model dtos.RoleUpdateDTO) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]dtos.RoleListDTO, error)
	FindByID(id uuid.UUID) (dtos.RoleListDTO, error)
	FindByUserID(id uuid.UUID) ([]dtos.RoleListDTO, error)
	CheckAdminByUserID(id uuid.UUID) (bool, error)
	CheckMemberByUserID(id uuid.UUID) (bool, error)
	FindByName(name string) (dtos.RoleListDTO, error)
}

//roleService is an implementation of RoleService
type roleService struct {
	roleRepository role.IRoleRepository
}

//NewRoleService is a constructor for RoleService
func NewRoleService(roleRepository role.IRoleRepository) RoleService {
	return &roleService{roleRepository: roleRepository}
}

//Create a new role
func (s *roleService) Create(model dtos.RoleCreateDTO) (dtos.RoleListDTO, error) {
	listModel := dtos.RoleListDTO{}
	roleModel := role.Role{
		Name:      model.Name,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := s.roleRepository.Create(&roleModel)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.RoleListDTO{
		ID:   roleModel.ID,
		Name: roleModel.Name,
	}
	return listModel, nil
}

//Update a role
func (s *roleService) Update(model dtos.RoleUpdateDTO) (dtos.RoleListDTO, error) {
	listModel := dtos.RoleListDTO{}
	roleModel := role.Role{
		ID:        model.ID,
		Name:      model.Name,
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := s.roleRepository.Update(&roleModel)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.RoleListDTO{
		ID:   roleModel.ID,
		Name: roleModel.Name,
	}
	return listModel, nil
}

//Delete a role
func (s *roleService) Delete(model dtos.RoleUpdateDTO) error {
	roleModel, err := s.roleRepository.FindByID(model.ID)
	if err != nil {
		return err
	}
	err = s.roleRepository.Delete(roleModel)
	if err != nil {
		return err
	}
	return nil
}

//DeleteByID a role
func (s *roleService) DeleteByID(id uuid.UUID) error {
	_, err := s.roleRepository.FindByID(id)
	if err != nil {
		return err
	}
	err = s.roleRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}

//FindAll roles
func (s *roleService) FindAll() ([]dtos.RoleListDTO, error) {
	listModel := []dtos.RoleListDTO{}
	roles, err := s.roleRepository.FindAll()
	if err != nil {
		return listModel, err
	}
	for _, role := range roles {
		listModel = append(listModel, dtos.RoleListDTO{
			ID:   role.ID,
			Name: role.Name,
		})
	}
	return listModel, nil
}

//FindByID a role
func (s *roleService) FindByID(id uuid.UUID) (dtos.RoleListDTO, error) {
	listModel := dtos.RoleListDTO{}
	roleModel, err := s.roleRepository.FindByID(id)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.RoleListDTO{
		ID:   roleModel.ID,
		Name: roleModel.Name,
	}
	return listModel, nil
}

//FindByUserID a role
func (s *roleService) FindByUserID(id uuid.UUID) ([]dtos.RoleListDTO, error) {
	listModel := []dtos.RoleListDTO{}
	roleModel, err := s.roleRepository.FindByUserID(id)
	if err != nil {
		return listModel, err
	}
	for _, role := range *roleModel {
		listModel = append(listModel, dtos.RoleListDTO{
			ID:   role.ID,
			Name: role.Name,
		})
	}
	return listModel, nil
}

//FindByName a role
func (s *roleService) FindByName(name string) (dtos.RoleListDTO, error) {
	listModel := dtos.RoleListDTO{}
	roleModel, err := s.roleRepository.FindByName(name)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.RoleListDTO{
		ID:   roleModel.ID,
		Name: roleModel.Name,
	}
	return listModel, nil
}

//CheckAdminByUserID a role
func (s *roleService) CheckAdminByUserID(id uuid.UUID) (bool, error) {
	roleModel, err := s.roleRepository.FindByUserID(id)
	if err != nil {
		return false, err
	}
	for _, role := range *roleModel {
		if role.Name == "admin" {
			return true, nil
		}
	}
	return false, nil
}

//CheckMemberByUserID a role
func (s *roleService) CheckMemberByUserID(id uuid.UUID) (bool, error) {
	roleModel, err := s.roleRepository.FindByUserID(id)
	if err != nil {
		return false, err
	}
	for _, role := range *roleModel {
		if role.Name == "member" {
			return true, nil
		}
	}
	return false, nil
}
