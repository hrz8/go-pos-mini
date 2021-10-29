package rest

import (
	"errors"
	"net/http"

	DomainUserError "github.com/hrz8/go-pos-mini/domains/user/error"
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
	if errors.Is(domainError, DomainUserError.Create.Err) {
		status := uint16(DomainUserError.Create.Status)
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
	if errors.Is(domainError, DomainUserError.Login.Err) {
		status := uint16(DomainUserError.Login.Status)
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
	if errors.Is(domainError, DomainUserError.UpdateById.Err) {
		status := uint16(DomainUserError.UpdateById.Status)
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
	if errors.Is(domainError, DomainUserError.DeleteById.Err) {
		status := uint16(DomainUserError.DeleteById.Status)
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
	if errors.Is(domainError, DomainUserError.GetById.Err) {
		status := uint16(DomainUserError.GetById.Status)
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
	if errors.Is(domainError, DomainUserError.GetAll.Err) {
		status := uint16(DomainUserError.GetAll.Status)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": dataErr.Error(),
			},
			domainError.Error(),
			status,
			i.prefix+"-006",
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
		prefix: "DOMAIN-USER",
	}
}
