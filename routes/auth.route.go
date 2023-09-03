package routes

import (
	"final-project-prakerja/controllers"
	// "final-project-prakerja/midleware"

	"github.com/labstack/echo/v4"
)




type AuthRoutesController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRoutesController {
	return AuthRoutesController{authController}
}

func (rc *AuthRoutesController) AuthRoute(e echo.Group) {
	router := e.Group("/auth")

	router.POST("/register", rc.authController.RegisterUser)
	router.POST("/login", rc.authController.SigningUser)
	router.GET("/refresh", rc.authController.RefreshAccessToken)
	router.GET("/logout", rc.authController.LogoutUser)
	// router.Use(midleware.DeserializeUser)
}