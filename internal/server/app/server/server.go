package server

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/cyril-jump/gophkeeper/internal/server/app/api"
	createnewblobdata "github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/create-new-blob-data"
	createnewcarddata "github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/create-new-card-data"
	createnewcreddata "github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/create-new-cred-data"
	createnewtextdata "github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/create-new-text-data"
	getallblobdata "github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/get-all-blob-data"
	getallcarddata "github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/get-all-card-data"
	getallcreddata "github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/get-all-cred-data"
	getalltextdata "github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/get-all-text-data"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/login"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/register"
	updateblobdatabyid "github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/update-blob-data-by-id"
	updatecarddatabyid "github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/update-card-data-by-id"
	updatecreddatabyid "github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/update-cred-data-by-id"
	updatetextdatabyid "github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/update-text-data-by-id"
	"github.com/cyril-jump/gophkeeper/internal/server/app/middlewares/cookie"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/auth"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/provider"
)

func Init(ctx context.Context, provider provider.Provider, auth auth.Strict) *echo.Echo {

	cookieMW := cookie.New(auth)

	createNewBlobDataReqs := createnewblobdata.Setup(ctx, provider)
	createNewCardDataReqs := createnewcarddata.Setup(ctx, provider)
	createNewCredDataReqs := createnewcreddata.Setup(ctx, provider)
	createNewTextDataReqs := createnewtextdata.Setup(ctx, provider)
	getAllBlobDataReqs := getallblobdata.Setup(ctx, provider)
	getAllCardDataReqs := getallcarddata.Setup(ctx, provider)
	getAllCredDataReqs := getallcreddata.Setup(ctx, provider)
	getAllTextDataReqs := getalltextdata.Setup(ctx, provider)
	updateBlobDataByIDReqs := updateblobdatabyid.Setup(ctx, provider)
	updateCardDataByIDReqs := updatecarddatabyid.Setup(ctx, provider)
	updateCredDataByIDReqs := updatecreddatabyid.Setup(ctx, provider)
	updateTextDataByIDReqs := updatetextdatabyid.Setup(ctx, provider)
	loginReqs := login.Setup(ctx, provider, auth)
	registerReqs := register.Setup(ctx, provider, auth)

	requests := joinRequests(
		createNewBlobDataReqs,
		createNewCardDataReqs,
		createNewCredDataReqs,
		createNewTextDataReqs,
		getAllBlobDataReqs,
		getAllCardDataReqs,
		getAllCredDataReqs,
		getAllTextDataReqs,
		updateBlobDataByIDReqs,
		updateCardDataByIDReqs,
		updateCredDataByIDReqs,
		updateTextDataByIDReqs,
		loginReqs,
		registerReqs,
	)

	e := echo.New()

	group := e.Group("")

	strictGroup := e.Group("")
	strictGroup.Use(cookieMW.SessionWithCookies)

	api.RegisterHandlers(group, requests)
	api.RegisterStrictHandlers(strictGroup, requests)

	return e

}
