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
	GetAllTextData(ctx context.Context, userID int) ([]domain.TextData, error)
}

func (r *Repo) GetAllTextDataDB(ctx context.Context, userID int) ([]domain.TextData, error) {

	return r.provider.GetAllTextData(ctx, userID)
}
