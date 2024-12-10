package repository

import (
	"20241209/domain"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (repo UserRepository) Create(user *domain.User) error {
	return repo.db.Create(&user).Error
}

func (repo UserRepository) Get(user *domain.User) error {
	return repo.db.Find(&user).Error
}

func (repo UserRepository) Update(user *domain.User) error {
	log.Println(user.Username)
	return repo.db.Where("username=?", user.Username).Update("failed_logins", user.FailedLogins).Error
}
