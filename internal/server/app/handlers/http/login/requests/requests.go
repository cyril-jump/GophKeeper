package requests

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/cyril-jump/gophkeeper/internal/server/app/domain"
)

type Strict interface {
	CreateCookie(c echo.Context, userID string) error
}

type Usecase interface {
	ProcessGetByCredentials(ctx context.Context, login, password string) (domain.User, error)
}

type Requests struct {
	ctx     context.Context
	usecase Usecase
	strict  Strict
}

func New(ctx context.Context, usecase Usecase, strict Strict) *Requests {
	return &Requests{
		ctx:     ctx,
		usecase: usecase,
		strict:  strict,
	}
}

func (r *Requests) Login(c echo.Context) error {

	var user domain.User

	body, err := io.ReadAll(c.Request().Body)
	if err != nil || len(body) == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if user.Login == "" || user.Password == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	if user, err = r.usecase.ProcessGetByCredentials(r.ctx, user.Login, user.Password); err != nil {
		c.NoContent(http.StatusInternalServerError)
	}

	if err = r.strict.CreateCookie(c, user.ID); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
