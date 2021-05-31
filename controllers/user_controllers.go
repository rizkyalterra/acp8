package controllers

import (
	"acp8/libs/database"
	"acp8/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetUserControllers(c echo.Context) error {

	data, err := database.GetUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": "Failed get data User",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "Success get data user",
		"data":    data,
	})
}

func RegisterControllers(c echo.Context) error {
	var userRegister models.UserRegister
	c.Bind(&userRegister)

	data, err := database.RegisterUser(userRegister)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": "Failed insert data User",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "Success insert data user",
		"data":    data,
	})
}

// func loginControllers(c echo.Context) error {
// 	email := c.FormValue("email")
// 	password := c.FormValue("password")

// 	user := User{1, email, password}
// 	return c.JSON(http.StatusOK, user)
// }
