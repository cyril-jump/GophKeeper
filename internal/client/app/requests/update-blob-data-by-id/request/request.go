package request

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/gommon/log"

	"github.com/cyril-jump/gophkeeper/internal/client/pkg/config"
	"github.com/cyril-jump/gophkeeper/internal/client/pkg/raw"
	rawtodomain "github.com/cyril-jump/gophkeeper/internal/client/pkg/raw-to-domain"
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

func (r *Request) UpdateBlobDataByID() error {

	file, err := ioutil.ReadFile(r.conf.Path)
	if err != nil {
		return err
	}
	var rawData raw.BlobData
	err = json.Unmarshal(file, &rawData)
	if err != nil {
		return err
	}

	data := rawtodomain.RawBlobToDomainBlob(rawData)

	resp, err := r.client.R().SetCookie(&http.Cookie{
		Name:  config.CookieKey.String(),
		Value: r.conf.CookieKey,
		Path:  config.CookiePath.String(),
	}).SetBody(data).Put(r.conf.ServerAddress + types.BlobPath.String())
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		log.Info("Data has been updated")
	case http.StatusBadRequest:
		log.Warn("Invalid request format")
	case http.StatusInternalServerError:
		log.Warn("Internal server error")
	case http.StatusUnauthorized:
		log.Warn("User not authenticated")
	}

	return nil
}
