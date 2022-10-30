package login

import (
	"context"

	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/login/adapters"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/login/requests"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/login/usecase"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/auth"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/provider"
)

func Setup(ctx context.Context, provider provider.Provider, strict auth.Strict) *requests.Requests {
	repo := adapters.New(provider)

	uc := usecase.New(repo)

	reqs := requests.New(ctx, uc, strict)

	return reqs
}
