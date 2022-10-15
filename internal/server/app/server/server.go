package server

import (
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
)

func Init() *echo.Echo {

	createNewBlobDataReqs := createnewblobdata.Setup()
	createNewCardDataReqs := createnewcarddata.Setup()
	createNewCredDataReqs := createnewcreddata.Setup()
	createNewTextDataReqs := createnewtextdata.Setup()
	getAllBlobDataReqs := getallblobdata.Setup()
	getAllCardDataReqs := getallcarddata.Setup()
	getAllCredDataReqs := getallcreddata.Setup()
	getAllTextDataReqs := getalltextdata.Setup()
	updateBlobDataByIDReqs := updateblobdatabyid.Setup()
	updateCardDataByIDReqs := updatecarddatabyid.Setup()
	updateCredDataByIDReqs := updatecreddatabyid.Setup()
	updateTextDataByIDReqs := updatetextdatabyid.Setup()
	loginReqs := login.Setup()
	registerReqs := register.Setup()

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

	g := e.Group("")

	api.RegisterHandlers(g, requests)

	return e

}
