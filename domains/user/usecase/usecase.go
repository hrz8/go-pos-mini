package usecase

import (
	"github.com/hrz8/go-pos-mini/domains/user/repository"
	"github.com/hrz8/go-pos-mini/models"
	"github.com/hrz8/go-pos-mini/utils"
	"golang.org/x/crypto/bcrypt"
)

type (
	UsecaseInterface interface {
		Create(ctx *utils.CustomContext, user *models.UserPayloadCreate) (*models.User, error)
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

func NewUsecase(repo repository.RepositoryInterface) UsecaseInterface {
	return &impl{
		repository: repo,
	}
}
