package routes

import (
	"20241209/infra"
	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	r.POST("/register", ctx.Ctl.User.Registration)
	r.POST("/login", ctx.Ctl.Auth.Login)

	return r
}
