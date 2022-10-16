package requests

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/cyril-jump/gophkeeper/internal/server/app/domain"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/config"
)

type Usecase interface {
	ProcessCreateNewBlobData(ctx context.Context, userID string, data domain.BlobData) error
}

type Requests struct {
	ctx     context.Context
	usecase Usecase
}

func New(ctx context.Context, usecase Usecase) *Requests {
	return &Requests{
		ctx:     ctx,
		usecase: usecase,
	}
}

func (r *Requests) CreateNewBlobData(c echo.Context) error {

	userID := ""

	if id := c.Request().Context().Value(config.CookieKey); id != nil {
		userID = id.(string)
	}

	file, err := c.FormFile("data")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return err
	}

	var inp domain.BlobData
	inp.Data = buf.Bytes()
	inp.Metadata = c.FormValue("metadata")
	if err := c.Bind(&inp); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = r.usecase.ProcessCreateNewBlobData(c.Request().Context(), userID, inp)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
