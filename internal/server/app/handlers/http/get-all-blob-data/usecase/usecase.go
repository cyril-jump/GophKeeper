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
	GetAllBlobDataDB(ctx context.Context, userID int) ([]domain.BlobData, error)
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessGetAllBlobData(ctx context.Context, userID int) ([]domain.BlobData, error) {
	log.Info("processing GetAllBlobData")

	data, err := u.repo.GetAllBlobDataDB(ctx, userID)
	if err != nil {
		return nil, err
	}

	return data, nil
}
