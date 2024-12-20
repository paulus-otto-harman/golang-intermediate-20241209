package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `json:"username"`
	Password     string `json:"password"`
	FailedLogins int    `gorm:"type:smallint" json:"-"`
}
