package request

import (
	"context"
	"encoding/json"
	"io/ioutil"
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

func (r *Request) Register() error {

	file, err := ioutil.ReadFile(r.conf.Path)
	if err != nil {
		return err
	}
	var data domain.User
	err = json.Unmarshal(file, &data)
	if err != nil {
		return err
	}

	resp, err := r.client.R().SetBody(data).Post(r.conf.ServerAddress + types.RegisterPath.String())
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		cookies := resp.Cookies()
		log.Info("Cookie: ", cookies[0].Value)
	case http.StatusBadRequest, http.StatusInternalServerError:
		log.Warn(resp.StatusCode())
	case http.StatusConflict:
		log.Info("Login is already occupied")
	}

	return nil
}
