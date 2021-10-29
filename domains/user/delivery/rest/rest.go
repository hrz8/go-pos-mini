package rest

import (
	"time"

	"github.com/golang-jwt/jwt"
	DomainUserError "github.com/hrz8/go-pos-mini/domains/user/error"
	"github.com/hrz8/go-pos-mini/domains/user/usecase"
	"github.com/hrz8/go-pos-mini/models"
	"github.com/hrz8/go-pos-mini/utils"
	"github.com/labstack/echo/v4"
)

type (
	RESTInterface interface {
		Create(c echo.Context) error
		Login(c echo.Context) error
		UpdateById(c echo.Context) error
		DeleteById(c echo.Context) error
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

func (i *impl) Login(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.UserPayloadLogin)

	// get user data from login usecase
	result, err := i.usecase.Login(ctx, payload)
	if err != nil {
		return i.restError.Throw(ctx, DomainUserError.Login.Err, err)
	}

	// create new jwt claims schema
	tokenTemp := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.UserJwt{
		ID: result.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
			Audience:  result.Email,
			Issuer:    "go-pos-mini-user-domain",
		},
	})
	token, err := tokenTemp.SignedString([]byte(ctx.AppConfig.SERVICE.JWTSECRET))
	if err != nil {
		return i.restError.Throw(ctx, DomainUserError.Login.Err, err)
	}

	return ctx.SuccessResponse(
		map[string]interface{}{"token": token},
		"success login",
		nil,
	)
}

func (i *impl) UpdateById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.UserPayloadUpdate)
	result, err := i.usecase.UpdateById(ctx, payload.ID, payload)
	if err != nil {
		return i.restError.Throw(ctx, DomainUserError.UpdateById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success update user",
		nil,
	)
}

func (i *impl) DeleteById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.UserPayloadDeleteById)
	result, err := i.usecase.DeleteById(ctx, payload.ID)
	if err != nil {
		return i.restError.Throw(ctx, DomainUserError.DeleteById.Err, err)
	}
	result.Password = nil
	return ctx.SuccessResponse(
		result,
		"success delete color",
		nil,
	)
}

func NewRest(u usecase.UsecaseInterface) RESTInterface {
	return &impl{
		usecase:   u,
		restError: NewPartnerError(),
	}
}
