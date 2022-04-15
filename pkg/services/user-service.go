package services

import (
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/user"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/helpers"
	"github.com/gofrs/uuid"
)

//UserService is an interface for UserService
type UserService interface {
	Create(model dtos.UserCreateDTO) (dtos.UserListDTO, error)
	Update(model dtos.UserUpdateDTO) (dtos.UserListDTO, error)
	Delete(model dtos.UserUpdateDTO) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]dtos.UserListDTO, error)
	FindByID(id uuid.UUID) (dtos.UserListDTO, error)
	FindByName(name string) (dtos.UserListDTO, error)
	FindByEmail(email string) (dtos.UserListDTO, error)
	FindByRoleId(roleId uuid.UUID) ([]dtos.UserListDTO, error)
	Search(ss string) ([]dtos.UserListDTO, error)
}

//userService is an implementation of UserService
type userService struct {
	userRepository user.IUserRepository
}

//NewUserService is a constructor for UserService
func NewUserService(userRepository user.IUserRepository) UserService {
	return &userService{userRepository: userRepository}
}

//Create a new user
func (s *userService) Create(model dtos.UserCreateDTO) (dtos.UserListDTO, error) {
	listModel := dtos.UserListDTO{}
	hashedPassword, _ := helpers.HashPassword(model.Password)
	userModel := user.User{
		Name:      model.Name,
		Email:     model.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		IsActive:  true, //TODO: default value should be true
	}
	err := s.userRepository.Create(&userModel)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.UserListDTO{
		ID:    userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
	}
	return listModel, nil
}

//Update a user
func (s *userService) Update(model dtos.UserUpdateDTO) (dtos.UserListDTO, error) {
	listModel := dtos.UserListDTO{}
	userModel, err := s.userRepository.FindByID(model.ID)
	if err != nil {
		return listModel, err
	}
	userModel.Name = model.Name
	userModel.Email = model.Email
	userModel.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err = s.userRepository.Update(userModel)
	if err != nil {
		return listModel, err
	}
	listModel = dtos.UserListDTO{
		ID:    model.ID,
		Name:  model.Name,
		Email: model.Email,
	}
	return listModel, nil
}

//Delete a user
func (s *userService) Delete(model dtos.UserUpdateDTO) error {
	userModel, err := s.userRepository.FindByID(model.ID)
	if err != nil {
		return err
	}
	err = s.userRepository.Delete(userModel)
	if err != nil {
		return err
	}
	return nil
}

//DeleteByID a user
func (s *userService) DeleteByID(id uuid.UUID) error {
	_, err := s.userRepository.FindByID(id)
	if err != nil {
		return err
	}
	err = s.userRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}

//FindAll users
func (s *userService) FindAll() ([]dtos.UserListDTO, error) {
	userModels, err := s.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	userList := []dtos.UserListDTO{}
	for _, userModel := range userModels {
		userList = append(userList, dtos.UserListDTO{
			ID:    userModel.ID,
			Name:  userModel.Name,
			Email: userModel.Email,
		})
	}
	return userList, nil
}

//FindByID a user
func (s *userService) FindByID(id uuid.UUID) (dtos.UserListDTO, error) {
	userModel, err := s.userRepository.FindByID(id)
	if err != nil {
		return dtos.UserListDTO{}, err
	}
	return dtos.UserListDTO{
		ID:    userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
	}, nil
}

//FindByName a user
func (s *userService) FindByName(name string) (dtos.UserListDTO, error) {
	userModel, err := s.userRepository.FindByName(name)
	if err != nil {
		return dtos.UserListDTO{}, err
	}
	return dtos.UserListDTO{
		ID:    userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
	}, nil
}

//FindByEmail a user
func (s *userService) FindByEmail(email string) (dtos.UserListDTO, error) {
	userModel, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return dtos.UserListDTO{}, err
	}
	return dtos.UserListDTO{
		ID:    userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
	}, nil
}

//FindByRoleId a user
func (s *userService) FindByRoleId(roleId uuid.UUID) ([]dtos.UserListDTO, error) {
	listModel := []dtos.UserListDTO{}
	userModels, err := s.userRepository.FindByRoleId(roleId)
	if err != nil {
		return nil, err
	}
	for _, userModel := range userModels {
		listModel = append(listModel, dtos.UserListDTO{
			ID:    userModel.ID,
			Name:  userModel.Name,
			Email: userModel.Email,
		})
	}
	return listModel, nil
}

//Search users
func (s *userService) Search(ss string) ([]dtos.UserListDTO, error) {
	userModels, err := s.userRepository.Search(ss)
	if err != nil {
		return nil, err
	}
	userList := []dtos.UserListDTO{}
	for _, userModel := range userModels {
		userList = append(userList, dtos.UserListDTO{
			ID:    userModel.ID,
			Name:  userModel.Name,
			Email: userModel.Email,
		})
	}
	return userList, nil
}
