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
	GetByCredentials(ctx context.Context, login, password string) (domain.User, error)
}

func (r *Repo) GetByCredentialsDB(ctx context.Context, login, password string) (domain.User, error) {

	return r.provider.GetByCredentials(ctx, login, password)
}
