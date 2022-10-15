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
	GetAllBlobData(ctx context.Context, userID int) ([]domain.BlobData, error)
}

func (r *Repo) GetAllBlobDataDB(ctx context.Context, userID int) ([]domain.BlobData, error) {

	return r.provider.GetAllBlobData(ctx, userID)
}
