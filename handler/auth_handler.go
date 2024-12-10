package handler

import (
	"20241209/domain"
	"20241209/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type AuthHandler struct {
	service service.AuthService
	logger  *zap.Logger
}

func NewAuthHandler(service service.AuthService, logger *zap.Logger) AuthHandler {
	return AuthHandler{service: service, logger: logger}
}

func (ctrl *AuthHandler) Login(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		BadResponse(c, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := ctrl.service.Login(user)
	if err != nil {
		BadResponse(c, "server error", http.StatusInternalServerError)
		return
	}

	if token == "" {
		BadResponse(c, "authentication failed", http.StatusUnauthorized)
		return
	}

	GoodResponseWithData(c, "user authenticated", http.StatusOK, token)
}
