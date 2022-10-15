package adapters

import (
	"context"

	"github.com/cyril-jump/gophkeeper/internal/server/app/domain"
)

type Repo struct {
	provider Provider
}

func New(p Provider) *Repo {
	return &Repo{
		provider: p,
	}
}

type Provider interface {
	CreateNewCredData(ctx context.Context, userID int, data domain.CredData) error
}

func (r *Repo) CreateNewCredDataDB(ctx context.Context, userID int, data domain.CredData) error {

	return r.provider.CreateNewCredData(ctx, userID, data)
}
