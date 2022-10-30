package requests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/cyril-jump/gophkeeper/internal/mocks"
	"github.com/cyril-jump/gophkeeper/internal/server/app/domain"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/update-cred-data-by-id/adapters"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/update-cred-data-by-id/usecase"
	"github.com/cyril-jump/gophkeeper/internal/server/app/middlewares/cookie"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/auth/strict"
	"github.com/cyril-jump/gophkeeper/internal/server/pkg/config"
)

type Suite struct {
	suite.Suite
	auth    *strict.Strict
	db      *mocks.MockProvider
	cfg     config.Config
	e       *echo.Echo
	ectx    *echo.Context
	router  *echo.Router
	testSrv *httptest.Server
	mw      *cookie.Cookie
	ctx     context.Context
	repo    *adapters.Repo
	reqs    *Requests
	uc      *usecase.Usecase
	ctrl    *gomock.Controller
}

func (suite *Suite) SetupTest() {
	suite.e = echo.New()
	suite.router = echo.NewRouter(suite.e)
	suite.ctrl = gomock.NewController(suite.T())
	suite.db = mocks.NewMockProvider(suite.ctrl)

	suite.cfg = config.Config{ServerAddress: ":9090"}
	suite.auth = strict.New(suite.ctx)

	suite.testSrv = httptest.NewServer(suite.e)

	suite.repo = adapters.New(suite.db)
	suite.uc = usecase.New(suite.repo)
	suite.reqs = New(suite.ctx, suite.uc)

	defer suite.ctrl.Finish()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) Test_UpdateCredDataByID_200() {
	suite.e.PUT("/", suite.reqs.UpdateCredDataByID)
	suite.db.EXPECT().UpdateCredDataByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	data := domain.CredData{
		ID:       0,
		Login:    "111",
		Password: "222",
		Metadata: "333",
	}

	reqBody, _ := json.Marshal(data)
	payload := strings.NewReader(string(reqBody))

	client := resty.New()
	res, err := client.R().SetBody(payload).Put(suite.testSrv.URL)

	if err != nil {
		log.Fatal("Could not create GET request")
	}
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode())

	defer suite.testSrv.Close()
}

func (suite *Suite) Test_UpdateCredDataByID_400() {
	suite.e.PUT("/", suite.reqs.UpdateCredDataByID)
	suite.db.EXPECT().UpdateCredDataByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	data := ""

	reqBody, _ := json.Marshal(data)
	payload := strings.NewReader(string(reqBody))

	client := resty.New()
	res, err := client.R().SetBody(payload).Put(suite.testSrv.URL)

	if err != nil {
		log.Fatal("Could not create GET request")
	}
	assert.Equal(suite.T(), http.StatusBadRequest, res.StatusCode())

	defer suite.testSrv.Close()
}

func (suite *Suite) Test_UpdateCredDataByID_401() {
	suite.e.Use(suite.mw.SessionWithCookies)
	suite.e.PUT("/", suite.reqs.UpdateCredDataByID)
	suite.db.EXPECT().UpdateCredDataByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	data := domain.CredData{
		ID:       0,
		Login:    "111",
		Password: "222",
		Metadata: "333",
	}

	reqBody, _ := json.Marshal(data)
	payload := strings.NewReader(string(reqBody))

	client := resty.New()
	res, err := client.R().SetBody(payload).Put(suite.testSrv.URL)

	if err != nil {
		log.Fatal("Could not create GET request")
	}
	assert.Equal(suite.T(), http.StatusUnauthorized, res.StatusCode())

	defer suite.testSrv.Close()
}

func (suite *Suite) Test_UpdateCredDataByID_500() {
	suite.e.PUT("/", suite.reqs.UpdateCredDataByID)
	suite.db.EXPECT().UpdateCredDataByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.ErrInternal)

	data := domain.CredData{
		ID:       0,
		Login:    "111",
		Password: "222",
		Metadata: "333",
	}

	reqBody, _ := json.Marshal(data)
	payload := strings.NewReader(string(reqBody))

	client := resty.New()
	res, err := client.R().SetBody(payload).Put(suite.testSrv.URL)

	if err != nil {
		log.Fatal("Could not create GET request")
	}
	assert.Equal(suite.T(), http.StatusInternalServerError, res.StatusCode())

	defer suite.testSrv.Close()
}
