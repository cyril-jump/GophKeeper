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
	GetAllCardDataDB(ctx context.Context, userID string) ([]domain.CardData, error)
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessGetAllCardData(ctx context.Context, userID string) ([]domain.CardData, error) {
	log.Info("processing GetAllCardData")

	data, err := u.repo.GetAllCardDataDB(ctx, userID)
	if err != nil {
		return nil, err
	}

	return data, nil
}
