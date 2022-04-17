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

// CartHandler interface for cart handler
type CartHandler interface {
	GetAllCarts(ctx *gin.Context)
	GetCart(ctx *gin.Context)
	GetCartsByUserID(ctx *gin.Context)
	GetCartsByUserIDInProgress(ctx *gin.Context)
	CreateCart(ctx *gin.Context)
	UpdateCart(ctx *gin.Context)
	DeleteCart(ctx *gin.Context)
	PassToOrder(ctx *gin.Context)
}

// cartHandler struct for cart handler
type cartHandler struct {
	cartService    services.CartService
	statusService  services.StatusService
	productService services.ProductService
	roleService    services.RoleService
	jwtService     services.JWTService
	orderService   services.OrderService
}

// NewCartHandler returns a new CartHandler
func NewCartHandler(cartService services.CartService, statusService services.StatusService, productService services.ProductService,
	roleService services.RoleService, jwtService services.JWTService, orderService services.OrderService) CartHandler {
	return &cartHandler{
		cartService:    cartService,
		statusService:  statusService,
		productService: productService,
		roleService:    roleService,
		jwtService:     jwtService,
		orderService:   orderService,
	}
}

// GetAllCarts returns all carts
func (h *cartHandler) GetAllCarts(ctx *gin.Context) {
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
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := helpers.StringToUUID(claims["user_id"].(string))
	//CheckMemberRole
	isMember, err := h.roleService.CheckMemberByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !isMember {
		response := helpers.BuildErrorResponse("Failed to process request", "You are not admin", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
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
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := helpers.StringToUUID(claims["user_id"].(string))
	//CheckMemberRole
	isMember, err := h.roleService.CheckMemberByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !isMember {
		response := helpers.BuildErrorResponse("Failed to process request", "You are not admin", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	//id := ctx.Param("id")
	//userID, _ := uuid.FromString(id)
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
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := helpers.StringToUUID(claims["user_id"].(string))
	//CheckMemberRole
	isMember, err := h.roleService.CheckMemberByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !isMember {
		response := helpers.BuildErrorResponse("Failed to process request", "You are not admin", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	inProgressId, err := h.statusService.FindByName("In Progress")
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request. Status Id is not valid.", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
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
	product, err := h.productService.FindByID(cart.ProductID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if &product == nil {
		response := helpers.BuildErrorResponse("Failed to process request", "Product is not valid", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	cart.Price = (float64(cart.Quantity)) * product.Price
	status, err := h.statusService.FindByName("In Progress")
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request. Status Id is not valid.", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	cart.StatusID = status.ID
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
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := helpers.StringToUUID(claims["user_id"].(string))
	//CheckMemberRole
	isMember, err := h.roleService.CheckMemberByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !isMember {
		response := helpers.BuildErrorResponse("Failed to process request", "You are not admin", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	var cart dtos.CartUpdateDTO
	err = ctx.BindJSON(&cart)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	cart.UserID = userID
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
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := helpers.StringToUUID(claims["user_id"].(string))
	//CheckMemberRole
	isMember, err := h.roleService.CheckMemberByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !isMember {
		response := helpers.BuildErrorResponse("Failed to process request", "You are not admin", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	id := ctx.Param("id")
	cartID, _ := uuid.FromString(id)
	IsOwner, err := h.cartService.CheckByUserIDAndID(userID, cartID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !IsOwner {
		response := helpers.BuildErrorResponse("Failed to process request", "You are not owner", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	err = h.cartService.DeleteByID(cartID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", helpers.EmptyResponse{})
	ctx.JSON(http.StatusNoContent, response)
}

//PassToOrder pass cart to order
func (h *cartHandler) PassToOrder(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := helpers.StringToUUID(claims["user_id"].(string))
	//CheckMemberRole
	isMember, err := h.roleService.CheckMemberByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !isMember {
		response := helpers.BuildErrorResponse("Failed to process request", "You are not admin", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	id := ctx.Param("id")
	cartID, _ := uuid.FromString(id)
	IsOwner, err := h.cartService.CheckByUserIDAndID(userID, cartID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !IsOwner {
		response := helpers.BuildErrorResponse("Failed to process request", "You are not owner", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	status, _ := h.statusService.FindByName("In Progress")
	statusCompleted, _ := h.statusService.FindByName("Completed")
	carts, err := h.cartService.FindByUserIDInProgress(userID, status.ID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	var orderCreateModels []dtos.OrderCreateDTO
	for _, cart := range carts {
		orderCreateModel := dtos.OrderCreateDTO{
			UserID:    userID,
			ProductID: cart.Product.ID,
			Quantity:  cart.Quantity,
			Price:     cart.Price,
			StatusID:  status.ID,
		}
		orderCreateModels = append(orderCreateModels, orderCreateModel)
	}
	//Create Order
	for _, orderCreateModel := range orderCreateModels {
		err = h.orderService.Create(orderCreateModel)
		if err != nil {
			response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}
	//Update Cart Status ID
	for _, cart := range carts {
		cartUpdateModel := dtos.CartUpdateDTO{
			ID:        cart.ID,
			UserID:    userID,
			ProductID: cart.Product.ID,
			Quantity:  cart.Quantity,
			StatusID:  statusCompleted.ID,
		}
		err = h.cartService.Update(cartUpdateModel)
		if err != nil {
			response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}
	response := helpers.BuildSuccessResponse(true, "Successful", helpers.EmptyResponse{})
	ctx.JSON(http.StatusNoContent, response)
}
