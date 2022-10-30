package requests

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/cyril-jump/gophkeeper/internal/server/app/domain"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/config"
)

type Usecase interface {
	ProcessGetAllTextData(ctx context.Context, userID string) ([]domain.TextData, error)
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

func (r *Requests) GetAllTextData(c echo.Context) error {

	userID := ""

	if id := c.Request().Context().Value(config.CookieKey); id != nil {
		userID = id.(string)
	}

	dataArray, err := r.usecase.ProcessGetAllTextData(c.Request().Context(), userID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrDataNotFound):
			return c.NoContent(http.StatusNoContent)
		default:
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	return c.JSON(http.StatusOK, dataArray)
}
