package midleware

import (
	"final-project-prakerja/initializers"
	"final-project-prakerja/models"
	"final-project-prakerja/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func DeserializeUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		
		var access_token string
		cookie, err := c.Cookie("access_token")

		authorizationHeader := c.Request().Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie.Value
		}
		fmt.Println(access_token)
		if access_token == "" {
			return c.JSON(http.StatusUnauthorized, "not logged in")
		}

		config, _ := initializers.LoadEnv()
		sub, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		var user models.User
		result := initializers.DB.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			
			return c.JSON(http.StatusForbidden, "the user belonging to this token no logger exists")

		}
		c.Set("currentUser", user)
	
		return next(c)
	}
}
