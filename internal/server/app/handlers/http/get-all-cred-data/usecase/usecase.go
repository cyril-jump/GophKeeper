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
	GetAllBlobCredDB(ctx context.Context, userID int) ([]domain.CredData, error)
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessGetAllCredData(ctx context.Context, userID int) ([]domain.CredData, error) {
	log.Info("processing GetAllCredData")

	data, err := u.repo.GetAllBlobCredDB(ctx, userID)
	if err != nil {
		return nil, err
	}

	return data, nil
}
