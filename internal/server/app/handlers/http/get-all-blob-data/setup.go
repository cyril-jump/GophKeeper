package getallblobdata

import (
	"context"

	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/get-all-blob-data/adapters"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/get-all-blob-data/requests"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/get-all-blob-data/usecase"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/provider"
)

func Setup(ctx context.Context, provider provider.Provider) *requests.Requests {
	repo := adapters.New(provider)

	uc := usecase.New(repo)

	reqs := requests.New(ctx, uc)

	return reqs
}
