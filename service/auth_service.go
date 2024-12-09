package service

import (
	"20241209/domain"
	"20241209/repository"
	"20241209/util"
)

type AuthService interface {
	Login(user domain.User) (bool, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(login domain.User) (bool, error) {
	user := domain.User{Username: login.Username}
	if err := s.repo.Get(&user); err != nil {
		return false, err
	}
	return util.CheckPasswordHash(login.Password, user.Password), nil
}
