package main

import (
	"final-project-prakerja/controllers"
	"final-project-prakerja/initializers"
	"final-project-prakerja/routes"
	"log"
	"net/http"

	// "github.com/gin-contrib/cors"
	"github.com/labstack/echo/v4"
	echomid "github.com/labstack/echo/v4/middleware"
)

var (
	server              *echo.Echo
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRoutesController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

func init() {
	config, err := initializers.LoadEnv()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables")
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	server = echo.New()

}

func main() {

	config, err := initializers.LoadEnv()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables")
	}

	corsConfig := echomid.CORSConfig{}
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true
	

	server.Use(echomid.CORSWithConfig(corsConfig))
	server.Use(echomid.LoggerWithConfig(echomid.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router := server.Group("/api")
	router.GET("/cek", func(c echo.Context) error {
		message := "welcome to golang with gorm and mysql"
		return c.JSON(http.StatusOK, message)

	})

	// routes.TestRoute(*router)
	AuthRouteController.AuthRoute(*router)
	UserRouteController.UserRoute(*router)

	server.Logger.Fatal(server.Start(":"+config.Port_server))

}
