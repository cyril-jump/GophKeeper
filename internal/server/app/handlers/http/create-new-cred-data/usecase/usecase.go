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
	CreateNewCredDataDB(ctx context.Context, userID int, data domain.CredData) error
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessCreateNewCredData(ctx context.Context, userID int, data domain.CredData) error {
	log.Info("processing CreateNewCredData")

	err := u.repo.CreateNewCredDataDB(ctx, userID, data)
	if err != nil {
		return err
	}

	return nil
}
