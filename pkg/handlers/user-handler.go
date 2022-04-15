package handlers

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/helpers"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

//UserHandler interface for user handler
type UserHandler interface {
	GetAllUsers(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetUserWithRoles(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

//userHandler struct for user handler
type userHandler struct {
	userService services.UserService
	roleService services.RoleService
}

//NewUserHandler returns a new UserHandler
func NewUserHandler(userService services.UserService, roleService services.RoleService) UserHandler {
	return &userHandler{
		userService: userService,
		roleService: roleService,
	}
}

//GetAllUsers returns all users
func (h *userHandler) GetAllUsers(ctx *gin.Context) {
	model := []dtos.UserListDTO{}
	users, err := h.userService.FindAll()
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	for _, v := range users {
		model = append(model, dtos.UserListDTO{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Email,
		})
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

//GetUser returns a user
func (h *userHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userUuid, _ := uuid.FromString(id)
	user, err := h.userService.FindByID(userUuid)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model := dtos.UserListDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

//GetUserWithRoles returns a user with roles
func (h *userHandler) GetUserWithRoles(ctx *gin.Context) {
	id := ctx.Param("id")
	userUuid, _ := uuid.FromString(id)
	user, err := h.userService.FindByID(userUuid)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	roles, err := h.roleService.FindByUserID(userUuid)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model := dtos.UserWithRolesDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Roles: roles,
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

//CreateUser creates a new user
func (h *userHandler) CreateUser(ctx *gin.Context) {
	var user dtos.UserCreateDTO
	err := ctx.BindJSON(&user)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	newUser, err := h.userService.Create(user)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", dtos.UserListDTO{
		ID:   newUser.ID,
		Name: newUser.Name,
	})
	ctx.JSON(http.StatusOK, response)
}

//UpdateUser updates a user
func (h *userHandler) UpdateUser(ctx *gin.Context) {
	var user dtos.UserUpdateDTO
	err := ctx.BindJSON(&user)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	updatedUser, err := h.userService.Update(user)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", dtos.UserListDTO{
		ID:    updatedUser.ID,
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
	})
	ctx.JSON(http.StatusOK, response)
}

//DeleteUser deletes a user
func (h *userHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userUuid, _ := uuid.FromString(id)
	err := h.userService.DeleteByID(userUuid)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", helpers.EmptyResponse{})
	ctx.JSON(http.StatusNoContent, response)
}
