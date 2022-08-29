package main

import (
	"GoProjedct/db"
	"GoProjedct/handler"
	"GoProjedct/log"
	"GoProjedct/repository/repo_impl"
	"GoProjedct/router"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func init() {
	os.Setenv("APP_NAME", "Github")
	log.InitLogger(false)
}

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

	log.Error("this is error")

	e := echo.New()
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3000"))
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome")
}
