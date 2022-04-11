package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthHandler interface for auth handler
type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

//authHandler struct for auth handler
type authHandler struct {
}

//NewAuthHandler returns a new AuthHandler
func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

func (h *authHandler) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK /*200*/, gin.H{
		"message": "Hello Login",
	})
}

func (h *authHandler) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK /*200*/, gin.H{
		"message": "Hello Register",
	})
}
