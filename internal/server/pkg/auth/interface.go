package auth

import "github.com/labstack/echo/v4"

type Strict interface {
	CreateToken(userID string) (string, error)
	CheckToken(tokenString string) (string, bool)
	CreateCookie(c echo.Context, userID string) error
}
