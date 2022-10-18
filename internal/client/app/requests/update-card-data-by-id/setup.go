package updatecarddatabyid

import (
	"context"

	"github.com/go-resty/resty/v2"

	"github.com/cyril-jump/gophkeeper/internal/client/app/requests/update-card-data-by-id/request"
	"github.com/cyril-jump/gophkeeper/internal/client/pkg/config"
)

func Setup(ctx context.Context, conf config.Config, client *resty.Client) *request.Request {

	reqs := request.New(ctx, conf, client)

	return reqs
}
