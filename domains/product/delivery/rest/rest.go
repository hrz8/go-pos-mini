package rest

import (
	DomainProductError "github.com/hrz8/go-pos-mini/domains/product/error"
	"github.com/hrz8/go-pos-mini/domains/product/usecase"
	"github.com/hrz8/go-pos-mini/models"
	"github.com/hrz8/go-pos-mini/utils"
	"github.com/labstack/echo/v4"
)

type (
	RESTInterface interface {
		Create(c echo.Context) error
		UpdateById(c echo.Context) error
		DeleteById(c echo.Context) error
		GetById(c echo.Context) error
		GetAll(c echo.Context) error
	}

	impl struct {
		usecase   usecase.UsecaseInterface
		restError RESTErrorInterface
	}
)

func (i *impl) Create(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProductPayloadCreate)
	result, err := i.usecase.Create(ctx, payload)
	if err != nil {
		return i.restError.Throw(ctx, DomainProductError.Create.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success create user",
		nil,
	)
}

func (i *impl) UpdateById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProductPayloadUpdate)
	result, err := i.usecase.UpdateById(ctx, payload.ID, payload)
	if err != nil {
		return i.restError.Throw(ctx, DomainProductError.UpdateById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success update user",
		nil,
	)
}

func (i *impl) DeleteById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProductPayloadDeleteById)
	result, err := i.usecase.DeleteById(ctx, payload.ID)
	if err != nil {
		return i.restError.Throw(ctx, DomainProductError.DeleteById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success delete user",
		nil,
	)
}

func (i *impl) GetById(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProductPayloadGetById)
	result, err := i.usecase.GetById(ctx, payload.ID)
	if err != nil {
		return i.restError.Throw(ctx, DomainProductError.GetById.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success get user by id",
		nil,
	)
}

func (i *impl) GetAll(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	payload := ctx.Payload.(*models.ProductPayloadGetAll)
	result, total, err := i.usecase.GetAll(ctx, payload)
	if err != nil {
		return i.restError.Throw(ctx, DomainProductError.GetAll.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success fetch all user",
		utils.ListMetaResponse{
			Count: len(*result),
			Total: int(*total),
		},
	)
}

func NewRest(u usecase.UsecaseInterface) RESTInterface {
	return &impl{
		usecase:   u,
		restError: NewPartnerError(),
	}
}
