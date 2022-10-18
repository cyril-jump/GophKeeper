package api

type ClientInterface interface {
	CreateNewBlobData() error

	CreateNewCardData() error

	CreateNewCredData() error

	CreateNewTextData() error

	GetAllBlobData() error

	GetAllCardData() error

	GetAllCredData() error

	GetAllTextData() error

	UpdateBlobDataByID() error

	UpdateCardDataByID() error

	UpdateCredDataByID() error

	UpdateTextDataByID() error

	Login() error

	Register() error
}
