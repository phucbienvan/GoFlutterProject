package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Phuc!")
	})
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")

	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")

	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
