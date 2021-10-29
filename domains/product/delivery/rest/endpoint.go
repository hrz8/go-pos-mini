package rest

import (
	"reflect"

	"github.com/hrz8/go-pos-mini/models"
	Utils "github.com/hrz8/go-pos-mini/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AddProductEndpoints(e *echo.Echo, rest RESTInterface, jwtMiddleware *middleware.JWTConfig) {
	e.POST("/api/v1/user", rest.Create, middleware.JWTWithConfig(*jwtMiddleware), Utils.ValidatorMiddleware(reflect.TypeOf(models.ProductPayloadCreate{})))
	e.PUT("/api/v1/user/:id", rest.UpdateById, middleware.JWTWithConfig(*jwtMiddleware), Utils.ValidatorMiddleware(reflect.TypeOf(models.ProductPayloadUpdate{})))
	e.DELETE("/api/v1/user/:id", rest.DeleteById, middleware.JWTWithConfig(*jwtMiddleware), Utils.ValidatorMiddleware(reflect.TypeOf(models.ProductPayloadDeleteById{})))
	e.GET("/api/v1/user/:id", rest.GetById, middleware.JWTWithConfig(*jwtMiddleware), Utils.ValidatorMiddleware(reflect.TypeOf(models.ProductPayloadGetById{})))
	e.GET("/api/v1/user", rest.GetAll, middleware.JWTWithConfig(*jwtMiddleware), Utils.ValidatorMiddleware(reflect.TypeOf(models.ProductPayloadGetAll{})))
}
