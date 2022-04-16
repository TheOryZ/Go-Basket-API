package handlers

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/helpers"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// ProductCategoryMapHandler interface for product category map handler
type ProductCategoryMapHandler interface {
	GetAllProductCategoryMaps(ctx *gin.Context)
	GetProductCategoryMap(ctx *gin.Context)
	CreateProductCategoryMap(ctx *gin.Context)
	UpdateProductCategoryMap(ctx *gin.Context)
	DeleteProductCategoryMap(ctx *gin.Context)
}

// productCategoryMapHandler struct for product category map handler
type productCategoryMapHandler struct {
	productCategoryMapService services.ProductCategoryMapService
	productService            services.ProductService
	categoryService           services.CategoryService
}

// NewProductCategoryMapHandler returns a new ProductCategoryMapHandler
func NewProductCategoryMapHandler(productCategoryMapService services.ProductCategoryMapService, productService services.ProductService, categoryService services.CategoryService) ProductCategoryMapHandler {
	return &productCategoryMapHandler{
		productCategoryMapService: productCategoryMapService,
		productService:            productService,
		categoryService:           categoryService,
	}
}

// GetAllProductCategoryMaps returns all product category maps
func (h *productCategoryMapHandler) GetAllProductCategoryMaps(ctx *gin.Context) {
	model := []dtos.ProductCategoryMapListDTO{}
	productCategoryMaps, err := h.productCategoryMapService.FindAll()
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	product, err := h.productService.FindByID(productCategoryMaps[0].ProductID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	category, err := h.categoryService.FindByID(productCategoryMaps[0].CategoryID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	productModel := dtos.ProductListDTO{
		ID:               product.ID,
		Name:             product.Name,
		SKU:              product.SKU,
		ShortDescription: product.ShortDescription,
		Description:      product.Description,
		Price:            product.Price,
		UnitOfStock:      product.UnitOfStock,
	}
	categoryModel := dtos.CategoryListDTO{
		ID:   category.ID,
		Name: category.Name,
	}
	for _, productCategoryMap := range productCategoryMaps {
		model = append(model, dtos.ProductCategoryMapListDTO{
			ID:       productCategoryMap.ID,
			Product:  &productModel,
			Category: &categoryModel,
		})
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// GetProductCategoryMap returns a product category map
func (h *productCategoryMapHandler) GetProductCategoryMap(ctx *gin.Context) {
	id := ctx.Param("id")
	_id, _ := uuid.FromString(id)
	productCategoryMap, err := h.productCategoryMapService.FindByID(_id)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", dtos.ProductCategoryMapListDTO{
		ID:       productCategoryMap.ID,
		Product:  &dtos.ProductListDTO{ID: productCategoryMap.Product.ID, Name: productCategoryMap.Product.Name},
		Category: &dtos.CategoryListDTO{ID: productCategoryMap.Category.ID, Name: productCategoryMap.Category.Name},
	})
	ctx.JSON(http.StatusOK, response)
}

// CreateProductCategoryMap creates a new product category map
func (h *productCategoryMapHandler) CreateProductCategoryMap(ctx *gin.Context) {
	var productCategoryMapCreateDTO dtos.ProductCategoryMapCreateDTO
	err := ctx.ShouldBindJSON(&productCategoryMapCreateDTO)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	product, err := h.productService.FindByID(productCategoryMapCreateDTO.ProductID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	category, err := h.categoryService.FindByID(productCategoryMapCreateDTO.CategoryID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	productCategoryMap := dtos.ProductCategoryMapCreateDTO{
		ProductID:  product.ID,
		CategoryID: category.ID,
	}
	model, err := h.productCategoryMapService.Create(productCategoryMap)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// UpdateProductCategoryMap updates a product category map
func (h *productCategoryMapHandler) UpdateProductCategoryMap(ctx *gin.Context) {
	var productCategoryMapUpdateDTO dtos.ProductCategoryMapUpdateDTO
	err := ctx.ShouldBindJSON(&productCategoryMapUpdateDTO)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	product, err := h.productService.FindByID(productCategoryMapUpdateDTO.ProductID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	category, err := h.categoryService.FindByID(productCategoryMapUpdateDTO.CategoryID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	productCategoryMap := dtos.ProductCategoryMapUpdateDTO{
		ProductID:  product.ID,
		CategoryID: category.ID,
	}
	model, err := h.productCategoryMapService.Update(productCategoryMap)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// DeleteProductCategoryMap deletes a product category map
func (h *productCategoryMapHandler) DeleteProductCategoryMap(ctx *gin.Context) {
	id := ctx.Param("id")
	_id, _ := uuid.FromString(id)
	err := h.productCategoryMapService.DeleteByID(_id)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", helpers.EmptyResponse{})
	ctx.JSON(http.StatusNoContent, response)
}
