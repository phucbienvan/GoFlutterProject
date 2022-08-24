package main

import (
	"GoProjedct/db"
	"GoProjedct/handle"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "",
		DbName:   "golang",
	}
	sql.Connect()

	defer sql.Close()

	e := echo.New()
	e.GET("/", welcome)
	e.GET("/users/sign-up", handle.SignUp)
	e.Logger.Fatal(e.Start(":3000"))
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome")
}
