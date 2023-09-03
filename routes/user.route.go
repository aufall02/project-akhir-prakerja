package routes

import (
	"final-project-prakerja/controllers"
	"final-project-prakerja/midleware"

	"github.com/labstack/echo/v4"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(e echo.Group){
	router := e.Group("/users", midleware.DeserializeUser)
	router.GET("/current", uc.userController.GetCurrent)
}