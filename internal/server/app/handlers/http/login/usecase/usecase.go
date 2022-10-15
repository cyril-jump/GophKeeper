package usecase

import (
	"context"

	"github.com/labstack/gommon/log"

	"github.com/cyril-jump/gophkeeper/internal/server/app/domain"
)

type Usecase struct {
	repo Repo
}

type Repo interface {
	GetByCredentialsDB(ctx context.Context, login, password string) (domain.User, error)
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessGetByCredentials(ctx context.Context, login, password string) (domain.User, error) {
	log.Info("processing GetByCredentials")

	user, err := u.repo.GetByCredentialsDB(ctx, login, password)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
