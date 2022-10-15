package requests

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Requests struct {
}

func New() *Requests {
	return &Requests{}
}

func (h *Requests) UpdateBlobDataByID(c echo.Context) error {

	return c.NoContent(http.StatusOK)
}