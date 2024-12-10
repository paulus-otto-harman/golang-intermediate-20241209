package service

import "20241209/repository"

type Service struct {
	Auth AuthService
	User UserService
}

func NewService(repo repository.Repository) Service {
	return Service{
		Auth: NewAuthService(repo),
		User: NewUserService(repo.User),
	}
}
