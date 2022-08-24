package handle

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SignUp(c echo.Context) error {
	type User struct {
		Email string
		Name  string
	}
	user := User{
		Email: "Phucbv@gmail.com",
		Name:  "Phuc",
	}
	return c.JSON(http.StatusOK, user)
}
