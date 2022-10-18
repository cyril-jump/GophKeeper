package request

import (
	"context"

	"github.com/go-resty/resty/v2"

	"github.com/cyril-jump/gophkeeper/internal/client/pkg/config"
)

type Request struct {
	ctx    context.Context
	conf   config.Config
	client *resty.Client
}

func New(ctx context.Context, conf config.Config, client *resty.Client) *Request {
	return &Request{
		ctx:    ctx,
		conf:   conf,
		client: client,
	}
}

func (r *Request) Login() error {

	return nil
}
