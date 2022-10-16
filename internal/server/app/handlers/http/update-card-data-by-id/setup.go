package updatecarddatabyid

import (
	"context"

	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/update-card-data-by-id/adapters"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/update-card-data-by-id/requests"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/update-card-data-by-id/usecase"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/provider"
)

func Setup(ctx context.Context, provider provider.Provider) *requests.Requests {
	repo := adapters.New(provider)

	uc := usecase.New(repo)

	reqs := requests.New(ctx, uc)

	return reqs
}
