package rest

import (
	"reflect"

	"github.com/hrz8/go-pos-mini/models"
	Utils "github.com/hrz8/go-pos-mini/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AddUserEndpoints(e *echo.Echo, rest RESTInterface, jwtMiddleware *middleware.JWTConfig) {
	e.POST("/api/v1/login", rest.Login, Utils.ValidatorMiddleware(reflect.TypeOf(models.UserPayloadLogin{})))
	e.POST("/api/v1/user", rest.Create, middleware.JWTWithConfig(*jwtMiddleware), Utils.ValidatorMiddleware(reflect.TypeOf(models.UserPayloadCreate{})))
	e.PUT("/api/v1/user/:id", rest.UpdateById, middleware.JWTWithConfig(*jwtMiddleware), Utils.ValidatorMiddleware(reflect.TypeOf(models.UserPayloadUpdate{})))
	e.DELETE("/api/v1/user/:id", rest.DeleteById, middleware.JWTWithConfig(*jwtMiddleware), Utils.ValidatorMiddleware(reflect.TypeOf(models.UserPayloadDeleteById{})))
}
