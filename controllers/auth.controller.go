package controllers

import (
	"final-project-prakerja/initializers"
	"final-project-prakerja/models"
	"final-project-prakerja/utils"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

// register user
func (ac *AuthController) RegisterUser(c echo.Context) error {
	var payload *models.UserRegisterRequest

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if payload.Password != payload.PasswordConfirm {
		return echo.NewHTTPError(http.StatusBadRequest, "Passwords do not match")
	}

	hashedPassword, err := utils.HashPassword(payload.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	now := time.Now()
	newUuid,_ := uuid.NewUUID()

	newUser := models.User{
		ID: newUuid,
		Username:  payload.Username,
		Name:      strings.ToLower(payload.Name),
		Password:  hashedPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return echo.NewHTTPError(http.StatusConflict, "User with that ID already exists")
	} else if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadGateway, "Something bad happened")
	}

	userResposne := models.UserResponse{
		ID:       newUser.ID.String(),
		Username: newUser.Username,
		Name:     newUser.Name,
	}

	return c.JSON(http.StatusCreated, userResposne)
}

// login user

func (ac *AuthController) SigningUser(c echo.Context) error {
	var payload models.UserLoginRequest

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var user models.User

	result := ac.DB.First(&user, "username = ?", payload.Username)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid username or Password")
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid username or Password")
	}

	config, _ := initializers.LoadEnv()
	ttlAcces, _ := time.ParseDuration(config.AccessTokenExpiresIn)
	ttlRefresh, _ := time.ParseDuration(config.RefreshTokenExpiresIn)
	maxAgeAcces, _ := time.ParseDuration(config.AccessTokenMaxAge)
	maxAgeRefresh, _ := time.ParseDuration(config.RefreshTokenMaxAge)

	

	//generate token
	access_token, err := utils.CreateToken(ttlAcces, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	refresh_token, err := utils.CreateToken(ttlRefresh, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	cookieAccess := http.Cookie{
		Name:     "access_token",
		Value:    access_token,
		Path:     "/",
		MaxAge:   int(maxAgeAcces) * 60,
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
	}
	cookieRefresh := http.Cookie{
		Name:     "refresh_token",
		Value:    refresh_token,
		Path:     "/",
		MaxAge:   int(maxAgeRefresh) * 60,
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
	}
	cookieLoggedIn := http.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   int(maxAgeAcces) * 60,
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: false,
	}

	c.SetCookie(&cookieAccess)
	c.SetCookie(&cookieRefresh)
	c.SetCookie(&cookieLoggedIn)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "succes",
		"acces_token": access_token,
	})
}

// Refresh Access Token
func (ac *AuthController) RefreshAccessToken(c echo.Context) error {
	message := "could not refress access_token"

	cookie, err := c.Cookie("refresh_token")

	if err != nil {
		return c.JSON(http.StatusForbidden, message)
	}

	config, _ := initializers.LoadEnv()

	sub, err := utils.ValidateToken(cookie.Value, config.AccessTokenPublicKey)
	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}

	var user models.User

	result := ac.DB.First(&user, "id = ?", fmt.Sprint(sub))
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusForbidden, "the user belonging to this token no logger exists")
	}

	ttlAccess, _ := time.ParseDuration(config.AccessTokenExpiresIn)
	maxAgeAccess, _ := time.ParseDuration(config.AccessTokenMaxAge)

	access_token, err := utils.CreateToken(ttlAccess, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}

	cookieAcces := http.Cookie{
		Name:     "access_token",
		Value:    access_token,
		Path:     "/",
		MaxAge:   int(maxAgeAccess) * 60,
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
	}
	cookieLoggedIn := http.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   int(maxAgeAccess) * 60,
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: false,
	}

	c.SetCookie(&cookieAcces)
	c.SetCookie(&cookieLoggedIn)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "succes",
		"acces_token": access_token,
	})
}

func (ac *AuthController) LogoutUser(c echo.Context) error {
	c.SetCookie(&http.Cookie{Name: "access_token", Value: ""})
	c.SetCookie(&http.Cookie{Name: "refresh_token", Value: ""})
	c.SetCookie(&http.Cookie{Name: "logged_in", Value: ""})

	return c.JSON(http.StatusOK, "success")
}
