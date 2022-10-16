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
	GetAllCardData(ctx context.Context, userID string) ([]domain.CardData, error)
}

func (r *Repo) GetAllCardDataDB(ctx context.Context, userID string) ([]domain.CardData, error) {

	return r.provider.GetAllCardData(ctx, userID)
}
