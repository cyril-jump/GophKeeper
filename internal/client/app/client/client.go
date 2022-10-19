package client

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/gommon/log"

	"github.com/cyril-jump/gophkeeper/internal/client/app/api"
	createnewblobdata "github.com/cyril-jump/gophkeeper/internal/client/app/requests/create-new-blob-data"
	createnewcarddata "github.com/cyril-jump/gophkeeper/internal/client/app/requests/create-new-card-data"
	createnewcreddata "github.com/cyril-jump/gophkeeper/internal/client/app/requests/create-new-cred-data"
	createnewtextdata "github.com/cyril-jump/gophkeeper/internal/client/app/requests/create-new-text-data"
	getallblobdata "github.com/cyril-jump/gophkeeper/internal/client/app/requests/get-all-blob-data"
	getallcarddata "github.com/cyril-jump/gophkeeper/internal/client/app/requests/get-all-card-data"
	getallcreddata "github.com/cyril-jump/gophkeeper/internal/client/app/requests/get-all-cred-data"
	getalltextdata "github.com/cyril-jump/gophkeeper/internal/client/app/requests/get-all-text-data"
	"github.com/cyril-jump/gophkeeper/internal/client/app/requests/login"
	"github.com/cyril-jump/gophkeeper/internal/client/app/requests/register"
	updateblobdatabyid "github.com/cyril-jump/gophkeeper/internal/client/app/requests/update-blob-data-by-id"
	updatecarddatabyid "github.com/cyril-jump/gophkeeper/internal/client/app/requests/update-card-data-by-id"
	updatecreddatabyid "github.com/cyril-jump/gophkeeper/internal/client/app/requests/update-cred-data-by-id"
	updatetextdatabyid "github.com/cyril-jump/gophkeeper/internal/client/app/requests/update-text-data-by-id"
	"github.com/cyril-jump/gophkeeper/internal/client/pkg/config"
	handlertype "github.com/cyril-jump/gophkeeper/internal/client/pkg/types"
)

// Client struct
type Client struct {
	Handler api.ClientInterface
}

// Init client
func Init(ctx context.Context, conf config.Config) *Client {

	client := resty.New()

	createNewBlobDataReqs := createnewblobdata.Setup(ctx, conf, client)
	createNewCardDataReqs := createnewcarddata.Setup(ctx, conf, client)
	createNewCredDataReqs := createnewcreddata.Setup(ctx, conf, client)
	createNewTextDataReqs := createnewtextdata.Setup(ctx, conf, client)
	getAllBlobDataReqs := getallblobdata.Setup(ctx, conf, client)
	getAllCardDataReqs := getallcarddata.Setup(ctx, conf, client)
	getAllCredDataReqs := getallcreddata.Setup(ctx, conf, client)
	getAllTextDataReqs := getalltextdata.Setup(ctx, conf, client)
	updateBlobDataByIDReqs := updateblobdatabyid.Setup(ctx, conf, client)
	updateCardDataByIDReqs := updatecarddatabyid.Setup(ctx, conf, client)
	updateCredDataByIDReqs := updatecreddatabyid.Setup(ctx, conf, client)
	updateTextDataByIDReqs := updatetextdatabyid.Setup(ctx, conf, client)
	loginReqs := login.Setup(ctx, conf, client)
	registerReqs := register.Setup(ctx, conf, client)

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

	return &Client{Handler: requests}
}

// Run client's method
func (c *Client) Run(handlerType handlertype.HandlerType) error {
	log.Info("Run")
	switch handlerType {
	case handlertype.Login:
		err := c.Handler.Login()
		if err != nil {
			return err
		}
	case handlertype.Register:
		err := c.Handler.Register()
		if err != nil {
			return err
		}
	case handlertype.CreateNewBlobData:
		err := c.Handler.CreateNewBlobData()
		if err != nil {
			return err
		}
	case handlertype.CreateNewCardData:
		err := c.Handler.CreateNewCardData()
		if err != nil {
			return err
		}
	case handlertype.CreateNewCredData:
		err := c.Handler.CreateNewCredData()
		if err != nil {
			return err
		}
	case handlertype.CreateNewTextData:
		err := c.Handler.CreateNewTextData()
		if err != nil {
			return err
		}
	case handlertype.GetAllBlobData:
		err := c.Handler.GetAllBlobData()
		if err != nil {
			return err
		}
	case handlertype.GetAllCardData:
		err := c.Handler.GetAllCardData()
		if err != nil {
			return err
		}
	case handlertype.GetAllCredData:
		err := c.Handler.GetAllCredData()
		if err != nil {
			return err
		}
	case handlertype.GetAllTextData:
		err := c.Handler.GetAllTextData()
		if err != nil {
			return err
		}
	case handlertype.UpdateBlobDataByID:
		err := c.Handler.UpdateBlobDataByID()
		if err != nil {
			return err
		}
	case handlertype.UpdateCardDataByID:
		err := c.Handler.UpdateCardDataByID()
		if err != nil {
			return err
		}
	case handlertype.UpdateCredDataByID:
		err := c.Handler.UpdateCredDataByID()
		if err != nil {
			return err
		}
	case handlertype.UpdateTextDataByID:
		err := c.Handler.UpdateTextDataByID()
		if err != nil {
			return err
		}
	default:
		return nil
	}

	return nil
}
