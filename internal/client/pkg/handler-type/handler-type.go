package handlertype

type HandlerType string

const (
	CreateNewBlobData HandlerType = "create-new-blob-data"
	CreateNewCardData HandlerType = "create-new-card-data"
	CreateNewCredData HandlerType = "create-new-cred-data"
	CreateNewTextData HandlerType = "create-new-text-data"

	GetAllBlobData HandlerType = "get-all-blob-data"
	GetAllCardData HandlerType = "get-all-card-data"
	GetAllCredData HandlerType = "get-all-cred-data"
	GetAllTextData HandlerType = "get-all-text-data"

	UpdateBlobDataByID HandlerType = "update-blob-data-by-id"
	UpdateCardDataByID HandlerType = "update-card-data-by-id"
	UpdateCredDataByID HandlerType = "update-cred-data-by-id"
	UpdateTextDataByID HandlerType = "update-text-data-by-id"

	Login    HandlerType = "login"
	Register HandlerType = "register"
)
