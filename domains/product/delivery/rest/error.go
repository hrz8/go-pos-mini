package rest

import (
	"errors"
	"net/http"

	DomainProductError "github.com/hrz8/go-pos-mini/domains/product/error"
	"github.com/hrz8/go-pos-mini/utils"
)

type (
	RESTErrorInterface interface {
		Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error
	}

	restErrorImpl struct {
		prefix string
	}
)

func (i *restErrorImpl) Throw(ctx *utils.CustomContext, domainError error, dataErr error) error {
	if errors.Is(domainError, DomainProductError.Create.Err) {
		status := uint16(DomainProductError.Create.Status)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": dataErr.Error(),
			},
			domainError.Error(),
			status,
			i.prefix+"-001",
			nil,
		)
	}
	if errors.Is(domainError, DomainProductError.UpdateById.Err) {
		status := uint16(DomainProductError.UpdateById.Status)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": dataErr.Error(),
			},
			domainError.Error(),
			status,
			i.prefix+"-002",
			nil,
		)
	}
	if errors.Is(domainError, DomainProductError.DeleteById.Err) {
		status := uint16(DomainProductError.DeleteById.Status)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": dataErr.Error(),
			},
			domainError.Error(),
			status,
			i.prefix+"-003",
			nil,
		)
	}
	if errors.Is(domainError, DomainProductError.GetById.Err) {
		status := uint16(DomainProductError.GetById.Status)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": dataErr.Error(),
			},
			domainError.Error(),
			status,
			i.prefix+"-004",
			nil,
		)
	}
	if errors.Is(domainError, DomainProductError.GetAll.Err) {
		status := uint16(DomainProductError.GetAll.Status)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": dataErr.Error(),
			},
			domainError.Error(),
			status,
			i.prefix+"-005",
			nil,
		)
	}
	return ctx.ErrorResponse(
		nil,
		"Internal Server Error",
		http.StatusInternalServerError,
		i.prefix+"-REST-500",
		nil,
	)
}

func NewPartnerError() RESTErrorInterface {
	return &restErrorImpl{
		prefix: "DOMAIN-PRODUCT",
	}
}
