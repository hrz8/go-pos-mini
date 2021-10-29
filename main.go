package main

import (
	"fmt"
	"net/http"

	Config "github.com/hrz8/go-pos-mini/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	appConfig := Config.NewConfig()

	restServer := echo.New()

	// middlewares
	restServer.Use(middleware.Logger())
	restServer.Pre(middleware.RemoveTrailingSlash())

	restServer.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	restServer.Logger.Fatal(restServer.Start(fmt.Sprintf(":%d", appConfig.SERVICE.RESTPORT)))
}
