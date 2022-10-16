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
	GetAllCredData(ctx context.Context, userID string) ([]domain.CredData, error)
}

func (r *Repo) GetAllCredDataDB(ctx context.Context, userID string) ([]domain.CredData, error) {

	return r.provider.GetAllCredData(ctx, userID)
}
