package repository

import (
	"20241209/database"
)

type AuthRepository struct {
	cache database.Cacher
}

func NewAuthRepository(cache database.Cacher) AuthRepository {
	return AuthRepository{cache: cache}
}

func (repo AuthRepository) Store(username, token string) error {
	if err := repo.cache.Set(username, token); err != nil {
		return err
	}
	return nil
}
