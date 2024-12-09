package handler

import (
	"20241209/domain"
	"20241209/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UserHandler struct {
	service service.UserService
	logger  *zap.Logger
}

func NewUserHandler(service service.UserService, logger *zap.Logger) UserHandler {
	return UserHandler{service: service, logger: logger}
}

func (ctrl *UserHandler) Registration(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		BadResponse(c, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := ctrl.service.Register(&user); err != nil {
		BadResponse(c, "server error", http.StatusInternalServerError)
		return
	}

	GoodResponseWithData(c, "user registered", http.StatusCreated, nil)
}
