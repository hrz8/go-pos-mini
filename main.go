package main

import (
	"fmt"
	"net/http"

	UserREST "github.com/hrz8/go-pos-mini/domains/user/delivery/rest"
	UserRepository "github.com/hrz8/go-pos-mini/domains/user/repository"
	UserUsecase "github.com/hrz8/go-pos-mini/domains/user/usecase"

	Config "github.com/hrz8/go-pos-mini/config"
	Database "github.com/hrz8/go-pos-mini/database"
	"github.com/hrz8/go-pos-mini/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	appConfig := Config.NewConfig()
	mysql := Database.NewMysql(appConfig)
	mysqlSess := mysql.Connect()

	RESTServer := echo.New()
	RESTServer.Validator = utils.NewValidator()

	// middlewares
	RESTServer.Use(middleware.Logger())
	RESTServer.Pre(middleware.RemoveTrailingSlash())
	RESTServer.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cc := &utils.CustomContext{
				Context:   ctx,
				MysqlSess: mysqlSess,
				AppConfig: appConfig,
			}
			return next(cc)
		}
	})

	RESTServer.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// domains functions
	userRepository := UserRepository.NewRepository(mysqlSess)
	userUsecase := UserUsecase.NewUsecase(userRepository)
	userREST := UserREST.NewRest(userUsecase)

	// register REST endpoints
	UserREST.AddUserEndpoints(RESTServer, userREST)

	// serve REST
	RESTServer.Logger.Fatal(RESTServer.Start(fmt.Sprintf(":%d", appConfig.SERVICE.RESTPORT)))
}
