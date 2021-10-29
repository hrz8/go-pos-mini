package rest

import (
	"reflect"

	"github.com/hrz8/go-pos-mini/models"
	Utils "github.com/hrz8/go-pos-mini/utils"
	"github.com/labstack/echo/v4"
)

func AddUserEndpoints(e *echo.Echo, rest RESTInterface) {
	e.POST("/api/v1/user", rest.Create, Utils.ValidatorMiddleware(reflect.TypeOf(models.UserPayloadCreate{})))
}
