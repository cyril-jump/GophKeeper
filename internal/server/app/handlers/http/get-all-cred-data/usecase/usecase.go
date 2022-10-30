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
	GetAllCredDataDB(ctx context.Context, userID string) ([]domain.CredData, error)
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessGetAllCredData(ctx context.Context, userID string) ([]domain.CredData, error) {
	log.Info("processing GetAllCredData")

	data, err := u.repo.GetAllCredDataDB(ctx, userID)
	if err != nil {
		return nil, err
	}

	return data, nil
}
