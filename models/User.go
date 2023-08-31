package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:100;not null;unique"`
	Password string `gorm:"size:100;not null;"`
	Name     string `gorm:"size:100;not null;"`
}
