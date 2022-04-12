package services

import (
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/user"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/helpers"
	"github.com/gofrs/uuid"
)

//AuthService is an interface for AuthService
type AuthService interface {
	CreateUser(model dtos.UserCreateDTO) (user.User, error)
	FindByEmail(email string) (user.User, error)
	VerifyUser(email string, password string) interface{}
	IsDuplicateEmail(email string) bool
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
		Name:      model.Name,
		Email:     model.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		IsActive:  true, //TODO: default value should be true
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

//VerifyUser verifies a user
func (s *authService) VerifyUser(email string, password string) interface{} {
	userModel, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return false
	}
	err = helpers.ComparePasswords(userModel.Password, password)
	if err != nil {
		return false
	}
	return *userModel
}

//IsDuplicateEmail checks if the email is already in use
func (s *authService) IsDuplicateEmail(email string) bool {
	userModel, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return false
	}
	return userModel.ID != uuid.Nil || userModel.Email != ""
}
