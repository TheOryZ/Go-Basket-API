package handlers

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/helpers"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// CartHandler interface for cart handler
type CartHandler interface {
	GetAllCarts(ctx *gin.Context)
	GetCart(ctx *gin.Context)
	GetCartsByUserID(ctx *gin.Context)
	GetCartsByUserIDInProgress(ctx *gin.Context)
	CreateCart(ctx *gin.Context)
	UpdateCart(ctx *gin.Context)
	DeleteCart(ctx *gin.Context)
} //TODO: endpoints be updated

// cartHandler struct for cart handler
type cartHandler struct {
	cartService   services.CartService
	statusService services.StatusService
}

// NewCartHandler returns a new CartHandler
func NewCartHandler(cartService services.CartService, statusService services.StatusService) CartHandler {
	return &cartHandler{
		cartService:   cartService,
		statusService: statusService,
	}
}

// GetAllCarts returns all carts
func (h *cartHandler) GetAllCarts(ctx *gin.Context) {
	carts, err := h.cartService.FindAll()
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", carts)
	ctx.JSON(http.StatusOK, response)
}

// GetCart returns a cart
func (h *cartHandler) GetCart(ctx *gin.Context) {
	id := ctx.Param("id")
	cartID, _ := uuid.FromString(id)
	cart, err := h.cartService.FindByID(cartID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", cart)
	ctx.JSON(http.StatusOK, response)
}

// GetCartsByUserID returns all carts by user id
func (h *cartHandler) GetCartsByUserID(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, _ := uuid.FromString(id)
	carts, err := h.cartService.FindByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", carts)
	ctx.JSON(http.StatusOK, response)
}

// GetCartsByUserIDInProgress returns all carts by user id in progress
func (h *cartHandler) GetCartsByUserIDInProgress(ctx *gin.Context) {
	id := ctx.Param("id")
	inProgressId, err := h.statusService.FindByName("In Progress")
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request. Status Id is not valid.", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userID, _ := uuid.FromString(id)
	carts, err := h.cartService.FindByUserIDInProgress(userID, inProgressId.ID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", carts)
	ctx.JSON(http.StatusOK, response)
}

// CreateCart creates a cart
func (h *cartHandler) CreateCart(ctx *gin.Context) {
	var cart dtos.CartCreateDTO
	err := ctx.BindJSON(&cart)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	err = h.cartService.Create(cart)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", cart)
	ctx.JSON(http.StatusOK, response)
}

// UpdateCart updates a cart
func (h *cartHandler) UpdateCart(ctx *gin.Context) {
	var cart dtos.CartUpdateDTO
	err := ctx.BindJSON(&cart)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	err = h.cartService.Update(cart)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", cart)
	ctx.JSON(http.StatusOK, response)
}

// DeleteCart deletes a cart
func (h *cartHandler) DeleteCart(ctx *gin.Context) {
	id := ctx.Param("id")
	cartID, _ := uuid.FromString(id)
	err := h.cartService.DeleteByID(cartID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", helpers.EmptyResponse{})
	ctx.JSON(http.StatusNoContent, response)
}
