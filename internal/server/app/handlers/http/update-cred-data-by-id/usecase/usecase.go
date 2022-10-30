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
	UpdateCredDataByIDDB(ctx context.Context, userID string, data domain.CredData) error
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ProcessUpdateCredDataByID(ctx context.Context, userID string, data domain.CredData) error {
	log.Info("processing UpdateCredDataByID")

	err := u.repo.UpdateCredDataByIDDB(ctx, userID, data)
	if err != nil {
		return err
	}

	return nil
}
