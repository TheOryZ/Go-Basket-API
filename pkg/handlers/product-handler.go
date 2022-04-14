package handlers

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/helpers"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// ProductHandler interface for product handler
type ProductHandler interface {
	GetAllProducts(ctx *gin.Context)
	GetProduct(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}

// productHandler struct for product handler
type productHandler struct {
	productService services.ProductService
}

// NewProductHandler returns a new ProductHandler
func NewProductHandler(productService services.ProductService) ProductHandler {
	return &productHandler{
		productService: productService,
	}
}

// GetAllProducts returns all products
func (h *productHandler) GetAllProducts(ctx *gin.Context) {
	model := []dtos.ProductListDTO{}
	products, err := h.productService.FindAll()
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	for _, v := range products {
		model = append(model, dtos.ProductListDTO{
			ID:               v.ID,
			Name:             v.Name,
			SKU:              v.SKU,
			ShortDescription: v.ShortDescription,
			Description:      v.Description,
			Price:            v.Price,
			UnitOfStock:      v.UnitOfStock,
		})
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// GetProduct returns a product
func (h *productHandler) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	productID, _ := uuid.FromString(id)
	product, err := h.productService.FindByID(productID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model := dtos.ProductListDTO{
		ID:               product.ID,
		Name:             product.Name,
		SKU:              product.SKU,
		ShortDescription: product.ShortDescription,
		Description:      product.Description,
		Price:            product.Price,
		UnitOfStock:      product.UnitOfStock,
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// CreateProduct creates a product
func (h *productHandler) CreateProduct(ctx *gin.Context) {
	var product dtos.ProductCreateDTO
	if err := ctx.ShouldBindJSON(&product); err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	newProduct, err := h.productService.Create(product)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model := dtos.ProductListDTO{
		ID:               newProduct.ID,
		Name:             newProduct.Name,
		SKU:              newProduct.SKU,
		ShortDescription: newProduct.ShortDescription,
		Description:      newProduct.Description,
		Price:            newProduct.Price,
		UnitOfStock:      newProduct.UnitOfStock,
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// UpdateProduct updates a product
func (h *productHandler) UpdateProduct(ctx *gin.Context) {
	var product dtos.ProductUpdateDTO
	if err := ctx.ShouldBindJSON(&product); err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	updatedProduct, err := h.productService.Update(product)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model := dtos.ProductListDTO{
		ID:               updatedProduct.ID,
		Name:             updatedProduct.Name,
		SKU:              updatedProduct.SKU,
		ShortDescription: updatedProduct.ShortDescription,
		Description:      updatedProduct.Description,
		Price:            updatedProduct.Price,
		UnitOfStock:      updatedProduct.UnitOfStock,
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// DeleteProduct deletes a product
func (h *productHandler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	productID, _ := uuid.FromString(id)
	err := h.productService.DeleteByID(productID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", helpers.EmptyResponse{})
	ctx.JSON(http.StatusNoContent, response)
}
