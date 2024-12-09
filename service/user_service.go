package service

import (
	"20241209/domain"
	"20241209/repository"
	"20241209/util"
)

type UserService interface {
	Register(user *domain.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *domain.User) error {
	passwordHash, _ := util.HashPassword(user.Password)
	user.Password = passwordHash
	return s.repo.Create(user)
}
