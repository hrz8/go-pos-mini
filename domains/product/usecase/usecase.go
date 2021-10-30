package usecase

import (
	DomainProductError "github.com/hrz8/go-pos-mini/domains/user/error"

	OutletsProductsRepository "github.com/hrz8/go-pos-mini/domains/outlets_products/repository"
	"github.com/hrz8/go-pos-mini/domains/product/repository"
	ProductUtils "github.com/hrz8/go-pos-mini/domains/product/utils"
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
		repository                repository.RepositoryInterface
		outletsProductsRepository OutletsProductsRepository.RepositoryInterface
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

	if err != nil {
		return nil, err
	}

	prices, err := i.outletsProductsRepository.GetPricesProductId(nil, &[]uint64{result.ID})
	if err != nil {
		return nil, err
	}

	for _, outlet := range result.Outlets {
		outlet.Price = ProductUtils.GetPrice(prices, &outlet.ID, &result.ID)
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

	idArr := make([]uint64, 0, len(*result))
	for _, val := range *result {
		idArr = append(idArr, val.ID)
	}

	prices, err := i.outletsProductsRepository.GetPricesProductId(nil, &idArr)
	if err != nil {
		return nil, nil, err
	}

	for _, product := range *result {
		for _, outlet := range product.Outlets {
			outlet.Price = ProductUtils.GetPrice(prices, &outlet.ID, &product.ID)
		}
	}

	return result, total, err
}

func NewUsecase(
	repo repository.RepositoryInterface,
	outletsProductsRepo OutletsProductsRepository.RepositoryInterface,
) UsecaseInterface {
	return &impl{
		repository:                repo,
		outletsProductsRepository: outletsProductsRepo,
	}
}
