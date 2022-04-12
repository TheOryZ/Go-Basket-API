package handlers

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/internal/store/domain/user"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/dtos"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/helpers"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-TheOryZ/pkg/services"
	"github.com/gin-gonic/gin"
)

//AuthHandler interface for auth handler
type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

//authHandler struct for auth handler
type authHandler struct {
	authService services.AuthService
	jwtService  services.JWTService
}

//NewAuthHandler returns a new AuthHandler
func NewAuthHandler(authService services.AuthService, jwtService services.JWTService) AuthHandler {
	return &authHandler{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (h *authHandler) Login(ctx *gin.Context) {
	var loginDto dtos.LoginDTO
	errDto := ctx.ShouldBind(&loginDto)
	if errDto != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDto.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := h.authService.VerifyUser(loginDto.Email, loginDto.Password)
	if v, ok := authResult.(user.User); ok {
		generatedToken := h.jwtService.GenerateToken(v.ID.String())
		response := helpers.BuildSuccessResponse(true, "Login Successful", generatedToken)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helpers.BuildErrorResponse("Login Failed", "Invalid Email or Password", helpers.EmptyResponse{})
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (h *authHandler) Register(ctx *gin.Context) {
	var registerDto dtos.UserCreateDTO
	errDto := ctx.ShouldBind(&registerDto)
	if errDto != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDto.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if h.authService.IsDuplicateEmail(registerDto.Email) {
		response := helpers.BuildErrorResponse("Failed to process request", "Email already exists", helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		createdUser, err := h.authService.CreateUser(registerDto)
		if err != nil {
			response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyResponse{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		//TODO: Add role to user
		generatedToken := h.jwtService.GenerateToken(createdUser.ID.String())
		response := helpers.BuildSuccessResponse(true, "Registration Successful", generatedToken)
		ctx.JSON(http.StatusOK, response)
		return
	}

}
