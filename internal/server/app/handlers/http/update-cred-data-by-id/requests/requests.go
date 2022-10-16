package requests

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/cyril-jump/gophkeeper/internal/server/app/domain"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/config"
)

type Usecase interface {
	ProcessUpdateCredDataByID(ctx context.Context, userID string, data domain.CredData) error
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

func (r *Requests) UpdateCredDataByID(c echo.Context) error {

	userID := ""

	if id := c.Request().Context().Value(config.CookieKey); id != nil {
		userID = id.(string)
	}

	var inp domain.CredData
	if err := c.Bind(&inp); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	err := r.usecase.ProcessUpdateCredDataByID(c.Request().Context(), userID, inp)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
