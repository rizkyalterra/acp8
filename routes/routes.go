package routes

import (
	"acp8/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUserControllers)
	// e.POST("/login", loginControllers)
	e.POST("/register", controllers.RegisterControllers)
	return e
}
