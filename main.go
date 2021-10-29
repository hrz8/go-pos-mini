package main

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	MerchantRepository "github.com/hrz8/go-pos-mini/domains/merchant/repository"
	OutletRepository "github.com/hrz8/go-pos-mini/domains/outlet/repository"
	UserREST "github.com/hrz8/go-pos-mini/domains/user/delivery/rest"
	DomainUserError "github.com/hrz8/go-pos-mini/domains/user/error"
	UserRepository "github.com/hrz8/go-pos-mini/domains/user/repository"
	UserUsecase "github.com/hrz8/go-pos-mini/domains/user/usecase"
	"github.com/hrz8/go-pos-mini/models"

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

	// domains functions
	userRepository := UserRepository.NewRepository(mysqlSess, appConfig)
	userUsecase := UserUsecase.NewUsecase(userRepository)
	userREST := UserREST.NewRest(userUsecase)
	MerchantRepository.NewRepository(mysqlSess, appConfig)
	OutletRepository.NewRepository(mysqlSess, appConfig)

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
	// jwt middleware handler
	JWTConfig := middleware.JWTConfig{
		Claims:     &models.UserJwt{},
		SigningKey: []byte(appConfig.SERVICE.JWTSECRET),
		ParseTokenFunc: func(token string, c echo.Context) (interface{}, error) {
			ctx := c.(*utils.CustomContext)
			parsedJwt, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				if jwt.GetSigningMethod("HS256") != token.Method {
					return nil, errors.New("unexpected signing method")
				}
				return []byte(appConfig.SERVICE.JWTSECRET), nil
			})
			if err != nil || parsedJwt == nil {
				return nil, errors.New("invalid token")
			}
			userEmail := parsedJwt.Claims.(jwt.MapClaims)["aud"].(string)
			user, err := userRepository.GetBy(nil, &models.User{Email: userEmail})
			if err != nil {
				return nil, err
			}
			if user == nil {
				return nil, DomainUserError.GetBy.Err
			}
			ctx.User = user
			return user, nil
		},
	}

	// register REST endpoints
	UserREST.AddUserEndpoints(RESTServer, userREST, &JWTConfig)

	// serve REST
	RESTServer.Logger.Fatal(RESTServer.Start(fmt.Sprintf(":%d", appConfig.SERVICE.RESTPORT)))
}
