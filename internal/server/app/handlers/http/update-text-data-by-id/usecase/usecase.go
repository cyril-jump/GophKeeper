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
	UpdateTextDataByIDDB(ctx context.Context, userID int, data domain.TextData) error
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessUpdateTextDataByID(ctx context.Context, userID int, data domain.TextData) error {
	log.Info("processing UpdateBlobDataByID")

	err := u.repo.UpdateTextDataByIDDB(ctx, userID, data)
	if err != nil {
		return err
	}

	return nil
}
