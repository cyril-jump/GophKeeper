package requests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/cyril-jump/gophkeeper/internal/mocks"
	"github.com/cyril-jump/gophkeeper/internal/server/app/domain"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/get-all-blob-data/adapters"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/get-all-blob-data/usecase"
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

func (suite *Suite) Test_GetAllBlobData_200() {
	suite.e.GET("/", suite.reqs.GetAllBlobData)
	suite.db.EXPECT().GetAllBlobData(gomock.Any(), gomock.Any()).Return([]domain.BlobData{}, nil)

	client := resty.New()
	res, err := client.R().Get(suite.testSrv.URL)

	if err != nil {
		log.Fatal("Could not create GET request")
	}
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode())

	defer suite.testSrv.Close()
}

func (suite *Suite) Test_GetAllBlobData_204() {
	suite.e.GET("/", suite.reqs.GetAllBlobData)
	suite.db.EXPECT().GetAllBlobData(gomock.Any(), gomock.Any()).Return([]domain.BlobData{}, domain.ErrDataNotFound)

	client := resty.New()
	res, err := client.R().Get(suite.testSrv.URL)

	if err != nil {
		log.Fatal("Could not create GET request")
	}
	assert.Equal(suite.T(), http.StatusNoContent, res.StatusCode())

	defer suite.testSrv.Close()
}

func (suite *Suite) Test_GetAllBlobData_401() {
	suite.e.Use(suite.mw.SessionWithCookies)
	suite.e.POST("/", suite.reqs.GetAllBlobData)
	suite.db.EXPECT().GetAllBlobData(gomock.Any(), gomock.Any()).Return([]domain.BlobData{}, nil)

	client := resty.New()
	res, err := client.R().Get(suite.testSrv.URL)

	if err != nil {
		log.Fatal("Could not create GET request")
	}
	assert.Equal(suite.T(), http.StatusUnauthorized, res.StatusCode())

	defer suite.testSrv.Close()
}

func (suite *Suite) Test_GetAllBlobData_500() {
	suite.e.GET("/", suite.reqs.GetAllBlobData)
	suite.db.EXPECT().GetAllBlobData(gomock.Any(), gomock.Any()).Return([]domain.BlobData{}, domain.ErrInternal)

	client := resty.New()
	res, err := client.R().Get(suite.testSrv.URL)

	if err != nil {
		log.Fatal("Could not create GET request")
	}
	assert.Equal(suite.T(), http.StatusInternalServerError, res.StatusCode())

	defer suite.testSrv.Close()
}
