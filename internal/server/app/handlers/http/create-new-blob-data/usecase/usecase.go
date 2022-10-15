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
	CreateNewBlobDataDB(ctx context.Context, userID int, data domain.BlobData) error
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessCreateNewBlobData(ctx context.Context, userID int, data domain.BlobData) error {
	log.Info("processing CreateNewBlobData")

	err := u.repo.CreateNewBlobDataDB(ctx, userID, data)
	if err != nil {
		return err
	}

	return nil
}
