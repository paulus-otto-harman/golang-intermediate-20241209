package service

import (
	"20241209/domain"
	"20241209/repository"
	"20241209/util"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type AuthService interface {
	Login(user domain.User) (string, error)
}

type authService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(login domain.User) (string, error) {
	user := domain.User{Username: login.Username}
	if err := s.repo.User.Get(&user); err != nil {
		return "", err
	}

	if util.CheckPasswordHash(login.Password, user.Password) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": login.Username,
			"exp":      time.Now().Add(time.Second * 30).Unix(),
		})

		if token == nil {
			return "", errors.New("failed to generate token")
		}

		secretKey := []byte("my-secret-key")

		// Proses Signing Token
		signedToken, err := token.SignedString(secretKey)
		if err != nil {
			return "", errors.New("failed to sign token")
		}

		s.repo.Auth.Store(login.Username, signedToken)

		return signedToken, nil
	}
	user.FailedLogins += 1
	s.repo.User.Update(&user)
	return "", nil
}
