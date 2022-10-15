package api

import "github.com/labstack/echo/v4"

// ServerInterface represents all server handlers.
type ServerInterface interface {
	CreateNewBlobData(c echo.Context) error

	CreateNewCardData(c echo.Context) error

	CreateNewCredData(c echo.Context) error

	CreateNewTextData(c echo.Context) error

	GetAllBlobData(c echo.Context) error

	GetAllCardData(c echo.Context) error

	GetAllCredData(c echo.Context) error

	GetAllTextData(c echo.Context) error

	UpdateBlobDataByID(c echo.Context) error

	UpdateCardDataByID(c echo.Context) error

	UpdateCredDataByID(c echo.Context) error

	UpdateTextDataByID(c echo.Context) error

	Login(c echo.Context) error

	Register(c echo.Context) error
}

type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// CreateNewBlobData converts echo context to params.
func (w *ServerInterfaceWrapper) CreateNewBlobData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateNewBlobData(ctx)
	return err
}

// CreateNewCardData converts echo context to params.
func (w *ServerInterfaceWrapper) CreateNewCardData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateNewCardData(ctx)
	return err
}

// CreateNewCredData converts echo context to params.
func (w *ServerInterfaceWrapper) CreateNewCredData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateNewCredData(ctx)
	return err
}

// CreateNewTextData converts echo context to params.
func (w *ServerInterfaceWrapper) CreateNewTextData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAllTextData(ctx)
	return err
}

// GetAllBlobData converts echo context to params.
func (w *ServerInterfaceWrapper) GetAllBlobData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAllBlobData(ctx)
	return err
}

// GetAllCardData converts echo context to params.
func (w *ServerInterfaceWrapper) GetAllCardData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAllCardData(ctx)
	return err
}

// GetAllCredData converts echo context to params.
func (w *ServerInterfaceWrapper) GetAllCredData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAllCredData(ctx)
	return err
}

// GetAllTextData converts echo context to params.
func (w *ServerInterfaceWrapper) GetAllTextData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAllTextData(ctx)
	return err
}

// UpdateBlobDataByID converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateBlobDataByID(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateBlobDataByID(ctx)
	return err
}

// UpdateCardDataByID converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateCardDataByID(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateCardDataByID(ctx)
	return err
}

// UpdateCredDataByID converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateCredDataByID(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateCredDataByID(ctx)
	return err
}

// UpdateTextDataByID converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateTextDataByID(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateTextDataByID(ctx)
	return err
}

// Login converts echo context to params.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Login(ctx)
	return err
}

// Register converts echo context to params.
func (w *ServerInterfaceWrapper) Register(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Register(ctx)
	return err
}

// RegisterHandlersWithBaseURL and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/api/login", wrapper.Login)
	router.POST(baseURL+"/api/register", wrapper.Register)
	router.POST(baseURL+"/api/materials/blob", wrapper.CreateNewBlobData)
	router.POST(baseURL+"/api/materials/card", wrapper.CreateNewCardData)
	router.POST(baseURL+"/api/materials/cred", wrapper.CreateNewCredData)
	router.POST(baseURL+"/api/materials/text", wrapper.CreateNewTextData)
	router.GET(baseURL+"/api/materials/blob", wrapper.GetAllBlobData)
	router.GET(baseURL+"/api/materials/card", wrapper.GetAllCardData)
	router.GET(baseURL+"/api/materials/cred", wrapper.GetAllCredData)
	router.GET(baseURL+"/api/materials/text", wrapper.GetAllTextData)
	router.PUT(baseURL+"/api/materials/blob", wrapper.UpdateBlobDataByID)
	router.PUT(baseURL+"/api/materials/card", wrapper.UpdateCardDataByID)
	router.PUT(baseURL+"/api/materials/cred", wrapper.UpdateCredDataByID)
	router.PUT(baseURL+"/api/materials/text", wrapper.UpdateTextDataByID)

}
