package usecase

import (
	DomainProductError "github.com/hrz8/go-pos-mini/domains/user/error"

	"github.com/hrz8/go-pos-mini/domains/product/repository"
	"github.com/hrz8/go-pos-mini/models"
	"github.com/hrz8/go-pos-mini/utils"
)

type (
	UsecaseInterface interface {
		Create(ctx *utils.CustomContext, user *models.ProductPayloadCreate) (*models.Product, error)
		UpdateById(ctx *utils.CustomContext, id uint64, payload *models.ProductPayloadUpdate) (*models.Product, error)
		DeleteById(ctx *utils.CustomContext, id uint64) (*models.Product, error)
		GetById(_ *utils.CustomContext, id uint64) (*models.Product, error)
		GetAll(_ *utils.CustomContext, conditions *models.ProductPayloadGetAll) (*[]models.Product, *int64, error)
	}

	impl struct {
		repository repository.RepositoryInterface
	}
)

func (i *impl) Create(ctx *utils.CustomContext, payload *models.ProductPayloadCreate) (*models.Product, error) {
	trx := ctx.MysqlSess.Begin()

	userCreated, err := i.repository.Create(trx, &models.Product{
		Name:        payload.Name,
		Description: payload.Description,
	})

	if err != nil {
		trx.Rollback()
		return nil, err
	}

	trx.Commit()
	return userCreated, err
}

func (i *impl) UpdateById(ctx *utils.CustomContext, id uint64, payload *models.ProductPayloadUpdate) (*models.Product, error) {
	trx := ctx.MysqlSess.Begin()
	instance, err := i.repository.GetBy(trx, &models.Product{ID: id})
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	result, err := i.repository.Update(trx, instance, payload)

	trx.Commit()
	return result, err
}

func (i *impl) DeleteById(ctx *utils.CustomContext, id uint64) (*models.Product, error) {
	trx := ctx.MysqlSess.Begin()
	instance, err := i.repository.GetBy(trx, &models.Product{ID: id})
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	if err := i.repository.DeleteById(nil, id); err != nil {
		trx.Rollback()
		return nil, err
	}

	trx.Commit()
	return instance, nil
}

func (i *impl) GetById(_ *utils.CustomContext, id uint64) (*models.Product, error) {
	result, err := i.repository.GetBy(nil, &models.Product{ID: id})
	if result == nil {
		return nil, DomainProductError.GetBy.Err
	}
	return result, err
}

func (i *impl) GetAll(_ *utils.CustomContext, conditions *models.ProductPayloadGetAll) (*[]models.Product, *int64, error) {
	result, err := i.repository.GetAll(nil, conditions)
	if err != nil {
		return nil, nil, err
	}
	total, err := i.repository.CountAll(nil)
	if err != nil {
		return nil, nil, err
	}
	return result, total, err
}

func NewUsecase(repo repository.RepositoryInterface) UsecaseInterface {
	return &impl{
		repository: repo,
	}
}
