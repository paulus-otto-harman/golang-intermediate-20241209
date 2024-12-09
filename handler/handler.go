package handler

import (
	"20241209/domain"
	"20241209/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	Auth AuthHandler
	User UserHandler
}

func NewHandler(service service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		Auth: NewAuthHandler(service.Auth, logger),
		User: NewUserHandler(service.User, logger),
	}
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func BadResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, Response{
		Status:  false,
		Message: message,
	})
}

func GoodResponseWithData(c *gin.Context, message string, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func GoodResponseWithPage(c *gin.Context, message string, statusCode, total, totalPages, page, Limit int, data interface{}) {
	c.JSON(statusCode, domain.DataPage{
		Status:      true,
		Message:     message,
		Total:       int64(total),
		Pages:       totalPages,
		CurrentPage: uint(page),
		Limit:       uint(Limit),
		Data:        data,
	})
}
