package strict

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/cyril-jump/gophkeeper/internal/server/pkg/config"
)

// Strict struct
type Strict struct {
	randNum []byte
	ctx     context.Context
}

// New Strict constructor
func New(ctx context.Context) *Strict {
	key := make([]byte, 16)

	_, err := rand.Read(key)
	if err != nil {
		log.Fatal(err)
	}

	return &Strict{
		randNum: key,
		ctx:     ctx,
	}
}

func (s *Strict) CreateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user": userID})
	tokenString, _ := token.SignedString(s.randNum)
	return tokenString, nil
}

func (s *Strict) CheckToken(tokenString string) (string, bool) {

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexected signing method: %v", token.Header["alg"])
		}
		return s.randNum, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return fmt.Sprintf("%s", claims["user"]), ok
	}
	return "", false
}

func (s *Strict) CreateCookie(c echo.Context, userID string) error {
	var err error
	cookie := new(http.Cookie)
	cookie.Path = "/"
	cookie.Value, err = s.CreateToken(userID)
	if err != nil {
		return err
	}
	cookie.Name = config.CookieKey.String()
	c.SetCookie(cookie)
	c.Request().AddCookie(cookie)
	return nil
}
