package usecase

import (
	DomainUserError "github.com/hrz8/go-pos-mini/domains/user/error"

	"github.com/hrz8/go-pos-mini/domains/user/repository"
	"github.com/hrz8/go-pos-mini/models"
	"github.com/hrz8/go-pos-mini/utils"
	"golang.org/x/crypto/bcrypt"
)

type (
	UsecaseInterface interface {
		Create(ctx *utils.CustomContext, user *models.UserPayloadCreate) (*models.User, error)
		Login(_ *utils.CustomContext, payload *models.UserPayloadLogin) (*models.User, error)
		UpdateById(ctx *utils.CustomContext, id uint64, payload *models.UserPayloadUpdate) (*models.User, error)
		DeleteById(ctx *utils.CustomContext, id uint64) (*models.User, error)
		GetById(_ *utils.CustomContext, id uint64) (*models.User, error)
		GetAll(_ *utils.CustomContext, conditions *models.UserPayloadGetAll) (*[]models.User, *int64, error)
	}

	impl struct {
		repository repository.RepositoryInterface
	}
)

func (i *impl) Create(ctx *utils.CustomContext, payload *models.UserPayloadCreate) (*models.User, error) {
	trx := ctx.MysqlSess.Begin()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	hashedPasswordStr := string(hashedPassword)
	userCreated, err := i.repository.Create(trx, &models.User{
		Email:     payload.Email,
		Password:  &hashedPasswordStr,
		FirstName: payload.FirstName,
		LastName:  &payload.LastName,
	})

	if err != nil {
		trx.Rollback()
		return nil, err
	}

	trx.Commit()
	return userCreated, err
}

func (i *impl) Login(_ *utils.CustomContext, payload *models.UserPayloadLogin) (*models.User, error) {
	result, err := i.repository.GetBy(nil, &models.User{Email: payload.Email})
	if result == nil {
		return nil, DomainUserError.GetBy.Err
	}

	wrongPassword := bcrypt.CompareHashAndPassword([]byte(*result.Password), []byte(payload.Password))

	if wrongPassword != nil {
		// login failed
		return nil, DomainUserError.WrongPassword.Err
	}

	return result, err
}

func (i *impl) UpdateById(ctx *utils.CustomContext, id uint64, payload *models.UserPayloadUpdate) (*models.User, error) {
	trx := ctx.MysqlSess.Begin()
	instance, err := i.repository.GetBy(trx, &models.User{ID: id})
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	if payload.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*payload.Password), bcrypt.DefaultCost)
		if err != nil {
			trx.Rollback()
			return nil, err
		}
		hashedPasswordStr := string(hashedPassword)
		payload.Password = &hashedPasswordStr
	}

	result, err := i.repository.Update(trx, instance, payload)

	trx.Commit()
	return result, err
}

func (i *impl) DeleteById(ctx *utils.CustomContext, id uint64) (*models.User, error) {
	trx := ctx.MysqlSess.Begin()
	instance, err := i.repository.GetBy(trx, &models.User{ID: id})
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

func (i *impl) GetById(_ *utils.CustomContext, id uint64) (*models.User, error) {
	result, err := i.repository.GetBy(nil, &models.User{ID: id})
	if result == nil {
		return nil, DomainUserError.GetBy.Err
	}
	return result, err
}

func (i *impl) GetAll(_ *utils.CustomContext, conditions *models.UserPayloadGetAll) (*[]models.User, *int64, error) {
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
