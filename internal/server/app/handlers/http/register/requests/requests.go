package requests

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/cyril-jump/gophkeeper/internal/server/app/domain"
)

type Usecase interface {
	ProcessCreate(ctx context.Context, user domain.User) (string, error)
}

type Strict interface {
	CreateCookie(c echo.Context, userID string) error
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

func (r *Requests) Register(c echo.Context) error {

	var user domain.User
	var userID string

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

	if userID, err = r.usecase.ProcessCreate(r.ctx, user); err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return c.NoContent(http.StatusConflict)
		}
		return c.NoContent(http.StatusInternalServerError)
	}

	if err = r.strict.CreateCookie(c, userID); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
