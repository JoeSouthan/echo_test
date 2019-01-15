package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// User - stores user info
type User struct {
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
}

// ErrorMessage - returns an error message
type ErrorMessage struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("echo", welcomeUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func welcomeUser(c echo.Context) (err error) {
	user := new(User)
	if err = c.Bind(user); err != nil {
		return c.JSON(http.StatusTeapot, &ErrorMessage{"User is not valid"})
	}
	return c.JSON(http.StatusOK, user)
}
