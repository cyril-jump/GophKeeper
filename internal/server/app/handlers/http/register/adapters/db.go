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
	Create(ctx context.Context, user domain.User) error
}

func (r *Repo) CreateDB(ctx context.Context, user domain.User) error {

	return r.provider.Create(ctx, user)
}
