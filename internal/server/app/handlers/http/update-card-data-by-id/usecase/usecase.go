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
	UpdateCardDataByIDDB(ctx context.Context, userID int, data domain.CardData) error
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessUpdateCardDataByID(ctx context.Context, userID int, data domain.CardData) error {
	log.Info("processing UpdateCardDataByID")

	err := u.repo.UpdateCardDataByIDDB(ctx, userID, data)
	if err != nil {
		return err
	}

	return nil
}
