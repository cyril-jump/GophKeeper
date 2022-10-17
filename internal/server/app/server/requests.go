package server

import (
	"github.com/labstack/echo/v4"

	"github.com/cyril-jump/gophkeeper/internal/server/app/api"
)

// CreateNewBlobDataMethod interface
type CreateNewBlobDataMethod interface {
	CreateNewBlobData(c echo.Context) error
}

// CreateNewBlobDataMethod interface
type CreateNewCardDataMethod interface {
	CreateNewCardData(c echo.Context) error
}

// CreateNewCredDataMethod interface
type CreateNewCredDataMethod interface {
	CreateNewCredData(c echo.Context) error
}

// CreateNewTextDataMethod interface
type CreateNewTextDataMethod interface {
	CreateNewTextData(c echo.Context) error
}

// GetAllBlobDataMethod interface
type GetAllBlobDataMethod interface {
	GetAllBlobData(c echo.Context) error
}

// GetAllCardDataMethod interface
type GetAllCardDataMethod interface {
	GetAllCardData(c echo.Context) error
}

// GetAllCredDataMethod interface
type GetAllCredDataMethod interface {
	GetAllCredData(c echo.Context) error
}

// GetAllTextDataMethod interface
type GetAllTextDataMethod interface {
	GetAllTextData(c echo.Context) error
}

// UpdateBlobDataByIDMethod interface
type UpdateBlobDataByIDMethod interface {
	UpdateBlobDataByID(c echo.Context) error
}

// UpdateCardDataByIDMethod interface
type UpdateCardDataByIDMethod interface {
	UpdateCardDataByID(c echo.Context) error
}

// UpdateCredDataByIDMethod interface
type UpdateCredDataByIDMethod interface {
	UpdateCredDataByID(c echo.Context) error
}

// UpdateTextDataByIDMethod interface
type UpdateTextDataByIDMethod interface {
	UpdateTextDataByID(c echo.Context) error
}

// LoginMethod interface
type LoginMethod interface {
	Login(c echo.Context) error
}

// RegisterMethod interface
type RegisterMethod interface {
	Register(c echo.Context) error
}

// joinRequests function for all handlers
func joinRequests(
	createNewBlobDataMethod CreateNewBlobDataMethod,
	createNewCardDataMethod CreateNewCardDataMethod,
	createNewCredDataMethod CreateNewCredDataMethod,
	createNewTextDataMethod CreateNewTextDataMethod,
	getAllBlobDataMethod GetAllBlobDataMethod,
	getAllCardDataMethod GetAllCardDataMethod,
	getAllCredDataMethod GetAllCredDataMethod,
	getAllTextDataMethod GetAllTextDataMethod,
	updateBlobDataByIDMethod UpdateBlobDataByIDMethod,
	updateCardDataByIDMethod UpdateCardDataByIDMethod,
	updateCredDataByIDMethod UpdateCredDataByIDMethod,
	updateTextDataByIDMethod UpdateTextDataByIDMethod,
	loginMethod LoginMethod,
	registerMethod RegisterMethod,
) api.ServerInterface {
	return &struct {
		CreateNewBlobDataMethod
		CreateNewCardDataMethod
		CreateNewCredDataMethod
		CreateNewTextDataMethod
		GetAllBlobDataMethod
		GetAllCardDataMethod
		GetAllCredDataMethod
		GetAllTextDataMethod
		UpdateBlobDataByIDMethod
		UpdateCardDataByIDMethod
		UpdateCredDataByIDMethod
		UpdateTextDataByIDMethod
		LoginMethod
		RegisterMethod
	}{
		CreateNewBlobDataMethod:  createNewBlobDataMethod,
		CreateNewCardDataMethod:  createNewCardDataMethod,
		CreateNewCredDataMethod:  createNewCredDataMethod,
		CreateNewTextDataMethod:  createNewTextDataMethod,
		GetAllBlobDataMethod:     getAllBlobDataMethod,
		GetAllCardDataMethod:     getAllCardDataMethod,
		GetAllCredDataMethod:     getAllCredDataMethod,
		GetAllTextDataMethod:     getAllTextDataMethod,
		UpdateBlobDataByIDMethod: updateBlobDataByIDMethod,
		UpdateCardDataByIDMethod: updateCardDataByIDMethod,
		UpdateCredDataByIDMethod: updateCredDataByIDMethod,
		UpdateTextDataByIDMethod: updateTextDataByIDMethod,
		LoginMethod:              loginMethod,
		RegisterMethod:           registerMethod,
	}
}
