package services

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/user"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/helpers"
)

//AuthService is an interface for AuthService
type AuthService interface {
	CreateUser(model dtos.UserCreateDTO) (user.User, error)
	FindByEmail(email string) (user.User, error)
}

//authService is an implementation of AuthService
type authService struct {
	userRepository user.IUserRepository
}

//NewAuthService creates a new AuthService
func NewAuthService(userRepository user.IUserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

//CreateUser creates a new user
func (s *authService) CreateUser(model dtos.UserCreateDTO) (user.User, error) {
	hashedPassword, _ := helpers.HashPassword(model.Password)
	userModel := user.User{
		Name:     model.Name,
		Email:    model.Email,
		Password: hashedPassword,
		IsActive: true, //TODO: default value should be true
	}
	err := s.userRepository.Create(&userModel)
	if err != nil {
		return user.User{}, err
	}
	return userModel, nil
}

//FindByEmail finds a user by email
func (s *authService) FindByEmail(email string) (user.User, error) {
	userModel, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return user.User{}, err
	}
	return *userModel, nil
}
