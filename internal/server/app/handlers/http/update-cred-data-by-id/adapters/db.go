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
	UpdateCredDataByID(ctx context.Context, userID string, data domain.CredData) error
}

func (r *Repo) UpdateCredDataByIDDB(ctx context.Context, userID string, data domain.CredData) error {

	return r.provider.UpdateCredDataByID(ctx, userID, data)
}
