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

// StatusHandler interface for status handler
type StatusHandler interface {
	GetAllStatus(ctx *gin.Context)
	GetStatus(ctx *gin.Context)
	CreateStatus(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
	DeleteStatus(ctx *gin.Context)
}

// statusHandler struct for status handler
type statusHandler struct {
	statusService services.StatusService
	jwtService    services.JWTService
	roleService   services.RoleService
}

// NewStatusHandler returns a new StatusHandler
func NewStatusHandler(statusService services.StatusService, jwtService services.JWTService, roleService services.RoleService) StatusHandler {
	return &statusHandler{
		statusService: statusService,
		roleService:   roleService,
		jwtService:    jwtService,
	}
}

// GetAllStatus returns all status
func (h *statusHandler) GetAllStatus(ctx *gin.Context) {
	model := []dtos.StatusListDTO{}
	status, err := h.statusService.FindAll()
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	for _, v := range status {
		model = append(model, dtos.StatusListDTO{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// GetStatus returns a status
func (h *statusHandler) GetStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	statusID, _ := uuid.FromString(id)
	status, err := h.statusService.FindByID(statusID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model := dtos.StatusListDTO{
		ID:   status.ID,
		Name: status.Name,
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// CreateStatus creates a status
func (h *statusHandler) CreateStatus(ctx *gin.Context) {
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
	var status dtos.StatusCreateDTO
	err = ctx.ShouldBindJSON(&status)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	newStatus, err := h.statusService.Create(status)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model := dtos.StatusListDTO{
		ID:   newStatus.ID,
		Name: newStatus.Name,
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// UpdateStatus updates a status
func (h *statusHandler) UpdateStatus(ctx *gin.Context) {
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
	var status dtos.StatusUpdateDTO
	err = ctx.ShouldBindJSON(&status)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	updatedStatus, err := h.statusService.Update(status)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", dtos.StatusListDTO{
		ID:   updatedStatus.ID,
		Name: updatedStatus.Name,
	})
	ctx.JSON(http.StatusOK, response)
}

// DeleteStatus deletes a status
func (h *statusHandler) DeleteStatus(ctx *gin.Context) {
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
	statusID, _ := uuid.FromString(id)
	err = h.statusService.DeleteByID(statusID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", helpers.EmptyResponse{})
	ctx.JSON(http.StatusNoContent, response)
}
