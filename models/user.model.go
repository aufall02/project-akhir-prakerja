package models

import (
	"github.com/google/uuid"
	"time"
)

// model user
type User struct {
	ID        uuid.UUID `gorm:"type:varchar(250);uuid_generate_v4();primarykey;unique"`
	Username  string    `gorm:"type:varchar(100);not null;unique"`
	Password  string    `gorm:"type:varchar(100);not null;"`
	Name      string    `gorm:"type:varchar(100);not null;"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

// user register request
type UserRegisterRequest struct {
	Username        string `json:"username" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required"`
}
// user login request
type UserLoginRequest struct {
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
}
// user response
type UserResponse struct {
	ID              string `json:"id" validate:"required"`
	Username        string `json:"username" validate:"required"`
	Name            string `json:"name" validate:"required"`
}
