package middleware

import (
	"20241209/handler"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Header struct {
	Token string `header:"Authorization" binding:"required"`
}

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := &Header{}
		if err := c.ShouldBindHeader(header); err != nil {
			handler.BadResponse(c, "invalid request", http.StatusBadRequest)
			c.Abort()
			return
		}
		//log.Println("header", header.Token)
		if err := checkToken(strings.Trim(strings.Replace(header.Token, "Bearer", "", 1), " ")); err != nil {
			handler.BadResponse(c, err.Error(), http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}

func checkToken(tokenString string) error {
	secretKey := []byte("my-secret-key")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifikasi algoritma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Memeriksa kadaluarsa token
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return errors.New("token is expired")
			}
		}
		log.Println("token is valid", claims)
		return nil
	}
	return errors.New("token is invalid")
}
