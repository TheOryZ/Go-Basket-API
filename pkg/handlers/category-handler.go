package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/helpers"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// CategoryHandler interface for category handler
type CategoryHandler interface {
	GetAllCategories(ctx *gin.Context)
	GetAllCategoriesPaging(ctx *gin.Context)
	GetCategory(ctx *gin.Context)
	GetCategoryWithProducts(ctx *gin.Context)
	UploadCsvFile(ctx *gin.Context)
	CreateCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
}

// categoryHandler struct for category handler
type categoryHandler struct {
	categoryService services.CategoryService
	productService  services.ProductService
	roleService     services.RoleService
	jwtService      services.JWTService
}

// NewCategoryHandler returns a new CategoryHandler
func NewCategoryHandler(categoryService services.CategoryService, productService services.ProductService, roleService services.RoleService, jwtService services.JWTService) CategoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
		productService:  productService,
		roleService:     roleService,
		jwtService:      jwtService,
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

//GetAllCategoriesPaging returns all categories with paging
func (h *categoryHandler) GetAllCategoriesPaging(ctx *gin.Context) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	pageInt, _ := helpers.StringToInt(page)
	limitInt, _ := helpers.StringToInt(limit)
	pageInt = (pageInt - 1) * limitInt
	categories, err := h.categoryService.FindAllWithPagination(pageInt, limitInt)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model := []dtos.CategoryListDTO{}
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

//GetCategoryWithProducts returns a category with products
func (h *categoryHandler) GetCategoryWithProducts(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryID, _ := uuid.FromString(id)
	category, err := h.categoryService.FindByID(categoryID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	products, err := h.productService.FindByCategoryID(categoryID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model := dtos.CategoryWithProductsDTO{
		ID:       category.ID,
		Name:     category.Name,
		Products: products,
	}
	response := helpers.BuildSuccessResponse(true, "Successful", model)
	ctx.JSON(http.StatusOK, response)
}

//UploadCsvFile uploads a csv file
func (h *categoryHandler) UploadCsvFile(ctx *gin.Context) {
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
	file, err := ctx.FormFile("file")
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	filename := filepath.Base(file.Filename)
	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	//ReadToCsv
	records, err := helpers.ReadToCsv(filename)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	//Create categories with records
	categoryListModel, err := h.categoryService.CreateAll(*records)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildSuccessResponse(true, "Successful", categoryListModel)
	ctx.JSON(http.StatusOK, response)
}

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
