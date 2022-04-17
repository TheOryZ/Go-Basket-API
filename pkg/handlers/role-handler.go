package handlers

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/helpers"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

//RoleHandler interface for role handler
type RoleHandler interface {
	GetAllRoles(ctx *gin.Context)
	GetRole(ctx *gin.Context)
	GetRoleWithUsers(ctx *gin.Context)
	CreateRole(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)
	DeleteRole(ctx *gin.Context)
}

//roleHandler struct for role handler
type roleHandler struct {
	roleService services.RoleService
	userService services.UserService
	jwtService  services.JWTService
}

//NewRoleHandler returns a new RoleHandler
func NewRoleHandler(roleService services.RoleService, userService services.UserService, jwtService services.JWTService) RoleHandler {
	return &roleHandler{
		roleService: roleService,
		userService: userService,
		jwtService:  jwtService,
	}
}

//GetAllRoles returns all roles
func (h *roleHandler) GetAllRoles(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := helpers.StringToUUID(claims["user_id"].(string))
	//CheckAdminRole
	isAdmin, err := h.roleService.CheckAdminByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !isAdmin {
		response := helpers.BuildErrorResponse("Failed to process request", "You are not admin", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	model := []dtos.RoleListDTO{}
	roles, err := h.roleService.FindAll()
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	for _, v := range roles {
		model = append(model, dtos.RoleListDTO{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

//GetRole returns a role
func (h *roleHandler) GetRole(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := helpers.StringToUUID(claims["user_id"].(string))
	//CheckAdminRole
	isAdmin, err := h.roleService.CheckAdminByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !isAdmin {
		response := helpers.BuildErrorResponse("Failed to process request", "You are not admin", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	id := ctx.Param("id")
	roleUuid, _ := uuid.FromString(id)
	role, err := h.roleService.FindByID(roleUuid)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", dtos.RoleListDTO{
		ID:   role.ID,
		Name: role.Name,
	})
	ctx.JSON(http.StatusOK, response)
}

//GetRoleWithUsers returns a role with users
func (h *roleHandler) GetRoleWithUsers(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := helpers.StringToUUID(claims["user_id"].(string))
	//CheckAdminRole
	isAdmin, err := h.roleService.CheckAdminByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !isAdmin {
		response := helpers.BuildErrorResponse("Failed to process request", "You are not admin", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	id := ctx.Param("id")
	roleUuid, _ := uuid.FromString(id)
	role, err := h.roleService.FindByID(roleUuid)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	users, err := h.userService.FindByRoleId(roleUuid)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model := dtos.RoleWithUserDTO{
		ID:   role.ID,
		Name: role.Name,
		User: users,
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

//CreateRole creates a new role
func (h *roleHandler) CreateRole(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := helpers.StringToUUID(claims["user_id"].(string))
	//CheckAdminRole
	isAdmin, err := h.roleService.CheckAdminByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !isAdmin {
		response := helpers.BuildErrorResponse("Failed to process request", "You are not admin", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	var role dtos.RoleCreateDTO
	err = ctx.BindJSON(&role)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	newRole, err := h.roleService.Create(role)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", dtos.RoleListDTO{
		ID:   newRole.ID,
		Name: newRole.Name,
	})
	ctx.JSON(http.StatusOK, response)
}

//UpdateRole updates a role
func (h *roleHandler) UpdateRole(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := helpers.StringToUUID(claims["user_id"].(string))
	//CheckAdminRole
	isAdmin, err := h.roleService.CheckAdminByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !isAdmin {
		response := helpers.BuildErrorResponse("Failed to process request", "You are not admin", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	var role dtos.RoleUpdateDTO
	err = ctx.BindJSON(&role)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	updatedRole, err := h.roleService.Update(role)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", dtos.RoleListDTO{
		ID:   updatedRole.ID,
		Name: updatedRole.Name,
	})
	ctx.JSON(http.StatusOK, response)
}

//DeleteRole deletes a role
func (h *roleHandler) DeleteRole(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := helpers.StringToUUID(claims["user_id"].(string))
	//CheckAdminRole
	isAdmin, err := h.roleService.CheckAdminByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !isAdmin {
		response := helpers.BuildErrorResponse("Failed to process request", "You are not admin", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	id := ctx.Param("id")
	roleUuid, _ := uuid.FromString(id)
	err = h.roleService.DeleteByID(roleUuid)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", helpers.EmptyResponse{})
	ctx.JSON(http.StatusNoContent, response)
}
