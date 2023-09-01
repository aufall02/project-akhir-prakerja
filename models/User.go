package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:varchar(250);uuid_generate_v4();primarykey;unique"`
	Username string    `gorm:"type:varchar(100);not null;unique"`
	Password string    `gorm:"type:varchar(100);not null;"`
	Name     string    `gorm:"type:varchar(100);not null;"`
}
