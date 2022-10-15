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
	CreateNewTextDataDB(ctx context.Context, userID int, data domain.TextData) error
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessCreateNewTextData(ctx context.Context, userID int, data domain.TextData) error {
	log.Info("processing CreateNewTextData")

	err := u.repo.CreateNewTextDataDB(ctx, userID, data)
	if err != nil {
		return err
	}

	return nil
}
