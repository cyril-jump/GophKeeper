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
	UpdateBlobDataByID(ctx context.Context, userID string, data domain.BlobData) error
}

func (r *Repo) UpdateBlobDataByIDDB(ctx context.Context, userID string, data domain.BlobData) error {

	return r.provider.UpdateBlobDataByID(ctx, userID, data)
}
