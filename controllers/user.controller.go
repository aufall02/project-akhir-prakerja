package controllers

import (
	"final-project-prakerja/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}


func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) GetCurrent(c echo.Context) error {
	currentUser := c.Get("currentUser").(models.User)

	userResponse := &models.UserResponse{
		ID: currentUser.ID.String(),
		Username: currentUser.Username,
		Name: currentUser.Name,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user" : userResponse,
	})
}
