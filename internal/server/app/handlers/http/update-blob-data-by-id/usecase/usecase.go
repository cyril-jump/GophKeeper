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
	UpdateBlobDataByIDDB(ctx context.Context, userID string, data domain.BlobData) error
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessUpdateBlobDataByID(ctx context.Context, userID string, data domain.BlobData) error {
	log.Info("processing UpdateBlobDataByID")

	err := u.repo.UpdateBlobDataByIDDB(ctx, userID, data)
	if err != nil {
		return err
	}

	return nil
}
