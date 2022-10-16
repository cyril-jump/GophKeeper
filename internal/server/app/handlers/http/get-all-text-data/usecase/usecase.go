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
	GetAllTextDataDB(ctx context.Context, userID string) ([]domain.TextData, error)
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessGetAllTextData(ctx context.Context, userID string) ([]domain.TextData, error) {
	log.Info("processing GetAllTextData")

	data, err := u.repo.GetAllTextDataDB(ctx, userID)
	if err != nil {
		return nil, err
	}

	return data, nil
}
