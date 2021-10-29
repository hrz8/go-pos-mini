package utils

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (c *CustomValidator) Validate(i interface{}) error {
	if err := c.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewValidator() echo.Validator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func ValidatorMiddleware(models reflect.Type) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(*CustomContext)
			payload := reflect.New(models).Interface()
			if err := ctx.Bind(payload); err != nil {
				if strings.Contains(err.Error(), "invalid syntax") {
					return ctx.ErrorResponse(
						nil,
						err.Error(),
						http.StatusBadRequest,
						"VALIDATOR-001",
						nil,
					)
				}
				return ctx.ErrorResponse(
					nil,
					"Internal Server Error",
					http.StatusInternalServerError,
					"VALIDATOR-500",
					nil,
				)
			}
			if err := ctx.Validate(payload); err != nil {
				return ctx.ErrorResponse(
					nil,
					err.Error(),
					http.StatusBadRequest,
					"VALIDATOR-002",
					nil,
				)
			}
			ctx.Payload = payload
			return next(ctx)
		}
	}
}
