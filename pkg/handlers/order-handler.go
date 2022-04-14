package handlers

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/helpers"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type OrderHandler interface {
	GetAllOrders(ctx *gin.Context)
	GetOrder(ctx *gin.Context)
	GetOrderByUser(ctx *gin.Context)
	//GetOrderByStatus(ctx *gin.Context) //TODO implement
	CreateOrder(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

type orderHandler struct {
	orderService   services.OrderService
	cartService    services.CartService
	productService services.ProductService
	userService    services.UserService
	statusService  services.StatusService
}

func NewOrderHandler(orderService services.OrderService, cartService services.CartService, productService services.ProductService, userService services.UserService, statusService services.StatusService) OrderHandler {
	return &orderHandler{
		orderService:   orderService,
		cartService:    cartService,
		productService: productService,
		userService:    userService,
		statusService:  statusService,
	}
}

// GetAllOrders returns all orders
func (h *orderHandler) GetAllOrders(ctx *gin.Context) {
	orders, err := h.orderService.FindAll()
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", orders)
	ctx.JSON(http.StatusOK, response)
}

// GetOrder returns a order
func (h *orderHandler) GetOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	orderID, _ := uuid.FromString(id)
	order, err := h.orderService.FindByID(orderID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", order)
	ctx.JSON(http.StatusOK, response)
}

// GetOrderByUser returns a order by user
func (h *orderHandler) GetOrderByUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, _ := uuid.FromString(id)
	order, err := h.orderService.FindByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", order)
	ctx.JSON(http.StatusOK, response)
}

// CreateOrder creates a order
func (h *orderHandler) CreateOrder(ctx *gin.Context) {
	var orderDTO dtos.OrderCreateDTO
	err := ctx.ShouldBindJSON(&orderDTO)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	err = h.orderService.Create(orderDTO)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userObj, _ := h.userService.FindByID(orderDTO.UserID)
	statusObj, _ := h.statusService.FindByID(orderDTO.StatusID)
	productObj, _ := h.productService.FindByID(orderDTO.ProductID)
	model := dtos.OrderListDTO{
		User:     userObj,
		Status:   statusObj,
		Product:  productObj,
		Quantity: orderDTO.Quantity,
		Price:    orderDTO.Price,
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// UpdateOrder updates a order
func (h *orderHandler) UpdateOrder(ctx *gin.Context) {
	var orderDTO dtos.OrderUpdateDTO
	err := ctx.ShouldBindJSON(&orderDTO)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	err = h.orderService.Update(orderDTO)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userObj, _ := h.userService.FindByID(orderDTO.UserID)
	statusObj, _ := h.statusService.FindByID(orderDTO.StatusID)
	productObj, _ := h.productService.FindByID(orderDTO.ProductID)
	model := dtos.OrderListDTO{
		User:     userObj,
		Status:   statusObj,
		Product:  productObj,
		Quantity: orderDTO.Quantity,
		Price:    orderDTO.Price,
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// DeleteOrder deletes a order
func (h *orderHandler) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	orderID, _ := uuid.FromString(id)
	err := h.orderService.DeleteByID(orderID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", helpers.EmptyResponse{})
	ctx.JSON(http.StatusOK, response)
}
