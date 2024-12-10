package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (m *Middleware) Whitelisted() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.RemoteIP() != "192.168.0.128" {
			log.Println("unauthorized access", c.RemoteIP())
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized access"})
			c.Abort()
			return
		}
		c.Next()
	}
}
