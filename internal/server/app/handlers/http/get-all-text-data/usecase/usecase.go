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
	CreateDB(ctx context.Context, user domain.User) error
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessSingUp(ctx context.Context, user domain.User) error {
	log.Info("processing SingUp")

	err := u.repo.CreateDB(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
