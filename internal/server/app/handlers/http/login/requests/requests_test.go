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
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/login/adapters"
	"github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/login/usecase"
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
	suite.reqs = New(suite.ctx, suite.uc, suite.auth)

	defer suite.ctrl.Finish()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) Test_Login_200() {
	suite.e.POST("/", suite.reqs.Login)
	suite.db.EXPECT().GetByCredentials(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.User{}, nil)
	user := domain.User{
		ID:       "",
		Login:    "111",
		Password: "222",
	}

	reqBody, _ := json.Marshal(user)
	payload := strings.NewReader(string(reqBody))

	client := resty.New()
	res, err := client.R().SetBody(payload).Post(suite.testSrv.URL)
	if err != nil {
		log.Fatal("Could not create POST request")
	}
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode())

	defer suite.testSrv.Close()
}

func (suite *Suite) Test_Login_400() {
	suite.e.POST("/", suite.reqs.Login)
	suite.db.EXPECT().GetByCredentials(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.User{}, nil)
	user := domain.User{
		ID:       "",
		Login:    "",
		Password: "",
	}

	reqBody, _ := json.Marshal(user)
	payload := strings.NewReader(string(reqBody))

	client := resty.New()
	res, err := client.R().SetBody(payload).Post(suite.testSrv.URL)
	if err != nil {
		log.Fatal("Could not create POST request")
	}
	assert.Equal(suite.T(), http.StatusBadRequest, res.StatusCode())

	defer suite.testSrv.Close()
}

func (suite *Suite) Test_Login_500() {
	suite.e.POST("/", suite.reqs.Login)
	suite.db.EXPECT().GetByCredentials(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.User{}, domain.ErrInternal)
	user := domain.User{
		ID:       "",
		Login:    "111",
		Password: "222",
	}

	reqBody, _ := json.Marshal(user)
	payload := strings.NewReader(string(reqBody))

	client := resty.New()
	res, err := client.R().SetBody(payload).Post(suite.testSrv.URL)
	if err != nil {
		log.Fatal("Could not create POST request")
	}
	assert.Equal(suite.T(), http.StatusInternalServerError, res.StatusCode())

	defer suite.testSrv.Close()
}
