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
	CreateNewCardData(ctx context.Context, userID int, data domain.CardData) error
}

func (r *Repo) CreateNewCardDataDB(ctx context.Context, userID int, data domain.CardData) error {

	return r.provider.CreateNewCardData(ctx, userID, data)
}
