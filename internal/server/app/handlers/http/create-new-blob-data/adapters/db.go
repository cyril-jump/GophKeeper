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
	CreateNewBlobData(ctx context.Context, userID int, data domain.BlobData) error
}

func (r *Repo) CreateNewBlobDataDB(ctx context.Context, userID int, data domain.BlobData) error {

	return r.provider.CreateNewBlobData(ctx, userID, data)
}
