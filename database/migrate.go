package database

import (
	"20241209/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	var err error

	if err = dropTables(db); err != nil {
		return err
	}

	if err = setupJoinTables(db); err != nil {
		return err
	}

	err = db.AutoMigrate(
		&domain.User{},
	)

	return err
}

func dropTables(db *gorm.DB) error {
	return db.Migrator().DropTable()
}

func setupJoinTables(db *gorm.DB) error {
	var err error

	return err
}
