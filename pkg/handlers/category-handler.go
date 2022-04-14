package handlers

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/helpers"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// CategoryHandler interface for category handler
type CategoryHandler interface {
	GetAllCategories(ctx *gin.Context)
	GetCategory(ctx *gin.Context)
	//CreateBulkCategories(ctx *gin.Context) //TODO: Implement this
	CreateCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
}

// categoryHandler struct for category handler
type categoryHandler struct {
	categoryService services.CategoryService
}

// NewCategoryHandler returns a new CategoryHandler
func NewCategoryHandler(categoryService services.CategoryService) CategoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}
}

// GetAllCategories returns all categories
func (h *categoryHandler) GetAllCategories(ctx *gin.Context) {
	model := []dtos.CategoryListDTO{}
	categories, err := h.categoryService.FindAll()
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	for _, v := range categories {
		model = append(model, dtos.CategoryListDTO{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// GetCategory returns a category
func (h *categoryHandler) GetCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryID, _ := uuid.FromString(id)
	category, err := h.categoryService.FindByID(categoryID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", dtos.CategoryListDTO{
		ID:   category.ID,
		Name: category.Name,
	})
	ctx.JSON(http.StatusOK, response)
}

// CreateBulkCategories creates a bulk of categories

// CreateCategory creates a category
func (h *categoryHandler) CreateCategory(ctx *gin.Context) {
	var category dtos.CategoryCreateDTO
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	newCategory, err := h.categoryService.Create(category)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model := dtos.CategoryListDTO{
		ID:   newCategory.ID,
		Name: newCategory.Name,
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// UpdateCategory updates a category
func (h *categoryHandler) UpdateCategory(ctx *gin.Context) {
	var category dtos.CategoryUpdateDTO
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	updatedCategory, err := h.categoryService.Update(category)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model := dtos.CategoryListDTO{
		ID:   updatedCategory.ID,
		Name: updatedCategory.Name,
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

// DeleteCategory deletes a category
func (h *categoryHandler) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryID, _ := uuid.FromString(id)
	err := h.categoryService.DeleteByID(categoryID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", helpers.EmptyResponse{})
	ctx.JSON(http.StatusNoContent, response)
}
