package server

import (
	"github.com/labstack/echo/v4"

	"github.com/cyril-jump/gophkeeper/internal/server/app/api"
)

type CreateNewBlobDataMethod interface {
	CreateNewBlobData(c echo.Context) error
}

type CreateNewCardDataMethod interface {
	CreateNewCardData(c echo.Context) error
}

type CreateNewCredDataMethod interface {
	CreateNewCredData(c echo.Context) error
}

type CreateNewTextDataMethod interface {
	CreateNewTextData(c echo.Context) error
}

type GetAllBlobDataMethod interface {
	GetAllBlobData(c echo.Context) error
}

type GetAllCardDataMethod interface {
	GetAllCardData(c echo.Context) error
}

type GetAllCredDataMethod interface {
	GetAllCredData(c echo.Context) error
}

type GetAllTextDataMethod interface {
	GetAllTextData(c echo.Context) error
}

type UpdateBlobDataByIDMethod interface {
	UpdateBlobDataByID(c echo.Context) error
}

type UpdateCardDataByIDMethod interface {
	UpdateCardDataByID(c echo.Context) error
}

type UpdateCredDataByIDMethod interface {
	UpdateCredDataByID(c echo.Context) error
}

type UpdateTextDataByIDMethod interface {
	UpdateTextDataByID(c echo.Context) error
}

type LoginMethod interface {
	Login(c echo.Context) error
}

type RegisterMethod interface {
	Register(c echo.Context) error
}

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
