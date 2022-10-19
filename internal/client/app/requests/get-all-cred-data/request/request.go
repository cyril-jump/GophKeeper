package request

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/gommon/log"

	"github.com/cyril-jump/gophkeeper/internal/client/app/domain"
	"github.com/cyril-jump/gophkeeper/internal/client/pkg/config"
	"github.com/cyril-jump/gophkeeper/internal/client/pkg/types"
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

func (r *Request) GetAllCredData() error {

	var data domain.CredData

	resp, err := r.client.R().SetCookie(&http.Cookie{
		Name:  config.CookieKey.String(),
		Value: r.conf.CookieKey,
		Path:  config.CookiePath.String(),
	}).Get(r.conf.ServerAddress + types.CredPath.String())
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		err = json.Unmarshal(resp.Body(), &data)
		if err != nil {
			return err
		}
		log.Info(data)
	case http.StatusNoContent:
		log.Warn("No response data")
	case http.StatusInternalServerError:
		log.Warn("Internal server error")
	case http.StatusUnauthorized:
		log.Warn("User not authenticated")
	}

	return nil
}
