package rest

import (
	DomainUserError "github.com/hrz8/go-pos-mini/domains/user/error"
	"github.com/hrz8/go-pos-mini/domains/user/usecase"
	"github.com/hrz8/go-pos-mini/models"
	"github.com/hrz8/go-pos-mini/utils"
	"github.com/labstack/echo/v4"
)

type (
	RESTInterface interface {
		Create(c echo.Context) error
	}

	impl struct {
		usecase   usecase.UsecaseInterface
		restError RESTErrorInterface
	}
)

func (i *impl) Create(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.UserPayloadCreate)
	result, err := i.usecase.Create(ctx, payload)
	if err != nil {
		return i.restError.Throw(ctx, DomainUserError.Create.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success create user",
		nil,
	)
}

func NewRest(u usecase.UsecaseInterface) RESTInterface {
	return &impl{
		usecase:   u,
		restError: NewPartnerError(),
	}
}
