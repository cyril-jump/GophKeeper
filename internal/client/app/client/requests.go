package client

import (
	"github.com/cyril-jump/gophkeeper/internal/client/app/api"
)

// CreateNewBlobDataMethod interface
type CreateNewBlobDataMethod interface {
	CreateNewBlobData() error
}

// CreateNewBlobDataMethod interface
type CreateNewCardDataMethod interface {
	CreateNewCardData() error
}

// CreateNewCredDataMethod interface
type CreateNewCredDataMethod interface {
	CreateNewCredData() error
}

// CreateNewTextDataMethod interface
type CreateNewTextDataMethod interface {
	CreateNewTextData() error
}

// GetAllBlobDataMethod interface
type GetAllBlobDataMethod interface {
	GetAllBlobData() error
}

// GetAllCardDataMethod interface
type GetAllCardDataMethod interface {
	GetAllCardData() error
}

// GetAllCredDataMethod interface
type GetAllCredDataMethod interface {
	GetAllCredData() error
}

// GetAllTextDataMethod interface
type GetAllTextDataMethod interface {
	GetAllTextData() error
}

// UpdateBlobDataByIDMethod interface
type UpdateBlobDataByIDMethod interface {
	UpdateBlobDataByID() error
}

// UpdateCardDataByIDMethod interface
type UpdateCardDataByIDMethod interface {
	UpdateCardDataByID() error
}

// UpdateCredDataByIDMethod interface
type UpdateCredDataByIDMethod interface {
	UpdateCredDataByID() error
}

// UpdateTextDataByIDMethod interface
type UpdateTextDataByIDMethod interface {
	UpdateTextDataByID() error
}

// LoginMethod interface
type LoginMethod interface {
	Login() error
}

// RegisterMethod interface
type RegisterMethod interface {
	Register() error
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
) api.ClientInterface {
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
