package register

import (
	"context"

	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/register/adapters"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/register/requests"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/register/usecase"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/auth"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/provider"
)

func Setup(ctx context.Context, provider provider.Provider, strict auth.Strict) *requests.Requests {
	repo := adapters.New(provider)

	uc := usecase.New(repo)

	reqs := requests.New(ctx, uc, strict)

	return reqs
}
