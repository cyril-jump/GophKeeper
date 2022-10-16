package cookie

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/cyril-jump/gophkeeper/internal/server/pkg/auth"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/config"
)

type Cookie struct {
	auth auth.Strict
}

func New(auth auth.Strict) *Cookie {
	return &Cookie{
		auth: auth,
	}
}

func (ck *Cookie) SessionWithCookies(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userID string
		var ok bool

		cookie, err := c.Cookie(config.CookieKey.String())
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		} else {
			userID, ok = ck.auth.CheckToken(cookie.Value)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
		}

		c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), config.CookieKey, userID)))

		return next(c)
	}
}
