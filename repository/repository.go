package repository

import (
	"20241209/config"
	"20241209/database"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Auth AuthRepository
	User UserRepository
}

func NewRepository(db *gorm.DB, cacher database.Cacher, config config.Config, log *zap.Logger) Repository {
	return Repository{
		Auth: NewAuthRepository(db, cacher, config.AppSecret),
		User: NewUserRepository(db),
	}
}
