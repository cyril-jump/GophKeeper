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
	CreateNewBlobCardDB(ctx context.Context, userID int, data domain.CardData) error
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessCreateNewCardData(ctx context.Context, userID int, data domain.CardData) error {
	log.Info("processing CreateNewCardData")

	err := u.repo.CreateNewBlobCardDB(ctx, userID, data)
	if err != nil {
		return err
	}

	return nil
}
