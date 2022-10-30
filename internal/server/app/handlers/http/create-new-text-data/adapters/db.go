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
	CreateNewTextData(ctx context.Context, userID string, data domain.TextData) error
}

func (r *Repo) CreateNewTextDataDB(ctx context.Context, userID string, data domain.TextData) error {

	return r.provider.CreateNewTextData(ctx, userID, data)
}
