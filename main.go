package main

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	MerchantRepository "github.com/hrz8/go-pos-mini/domains/merchant/repository"
	OutletRepository "github.com/hrz8/go-pos-mini/domains/outlet/repository"
	OutletsProductsRepository "github.com/hrz8/go-pos-mini/domains/outlets_products/repository"
	ProductREST "github.com/hrz8/go-pos-mini/domains/product/delivery/rest"
	ProductRepository "github.com/hrz8/go-pos-mini/domains/product/repository"
	ProductUsecase "github.com/hrz8/go-pos-mini/domains/product/usecase"
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
	userRepository := UserRepository.NewRepository(mysqlSess)
	userUsecase := UserUsecase.NewUsecase(userRepository)
	userREST := UserREST.NewRest(userUsecase)

	outletsProductsRepository := OutletsProductsRepository.NewRepository(mysqlSess)
	productRepository := ProductRepository.NewRepository(mysqlSess)
	productUsecase := ProductUsecase.NewUsecase(productRepository, outletsProductsRepository)
	productREST := ProductREST.NewRest(productUsecase)

	MerchantRepository.NewRepository(mysqlSess)
	OutletRepository.NewRepository(mysqlSess)

	// seeding
	UserRepository.RunSeed(mysqlSess, appConfig)
	ProductRepository.RunSeed(mysqlSess, appConfig)
	MerchantRepository.RunSeed(mysqlSess, appConfig)
	OutletRepository.RunSeed(mysqlSess, appConfig)

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
	ProductREST.AddProductEndpoints(RESTServer, productREST, &JWTConfig)

	// serve REST
	RESTServer.Logger.Fatal(RESTServer.Start(fmt.Sprintf(":%d", appConfig.SERVICE.RESTPORT)))
}
