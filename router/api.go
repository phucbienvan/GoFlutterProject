package router

import (
	"GoProjedct/handler"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/users/sign-up", api.UserHandler.SignUp)
}
