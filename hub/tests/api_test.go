package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"log/slog"

	"github.com/google/uuid"
	"github.com/qwark97/go-versatility-presentation/hub/peripherals/storage"
	"github.com/qwark97/go-versatility-presentation/hub/server"
	"github.com/qwark97/go-versatility-presentation/hub/server/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	peripheralsMock *mocks.Peripherals
	confMock        *mocks.Conf
	log             *slog.Logger
}

func TestSuite(t *testing.T) {
	s := Suite{
		peripheralsMock: mocks.NewPeripherals(t),
		confMock:        mocks.NewConf(t),
		log:             slog.Default(),
	}
	suite.Run(t, &s)
}

func (s Suite) SetupSuite() {
	s.confMock.EXPECT().Addr().Return("localhost:8080")
	s.confMock.EXPECT().RequestTimeout().Return(time.Second)
	serv := server.New(s.peripheralsMock, s.confMock, s.log)
	go func() {
		err := serv.Start()
		if err != nil {
			panic(err)
		}
	}()

	// wait for server to start
	time.Sleep(time.Second)
}

func (s Suite) TestShouldReceiveAllConfigurations() {
	assertion := assert.New(s.T())

	// given
	expectedConfigurations := []storage.Configuration{
		{
			ID: uuid.MustParse("eda1a0ca-12a3-4d3e-97e6-5baa0c5a1b93"),
		},
		{
			ID: uuid.MustParse("d19c786f-f77d-4ea7-bf83-6200d98e1402"),
		},
	}
	expectedStatusCode := http.StatusOK
	s.peripheralsMock.EXPECT().All(mock.Anything).Return(expectedConfigurations, nil).Once()
	client := http.Client{
		Timeout: time.Duration(time.Second),
	}
	apiPath := "http://localhost:8080/api/v1/configurations"

	// when
	resp, err := client.Get(apiPath)

	// then
	assertion.Nil(err)
	var actualResponse []storage.Configuration
	err = json.NewDecoder(resp.Body).Decode(&actualResponse)
	assertion.Nil(err)
	assertion.Equal(expectedStatusCode, resp.StatusCode)
	assertion.ElementsMatch(expectedConfigurations, actualResponse)
}

func (s Suite) TestShouldAddConfiguration() {
	assertion := assert.New(s.T())

	// given
	id := uuid.MustParse("eda1a0ca-12a3-4d3e-97e6-5baa0c5a1b93")
	expectedConfiguration := storage.Configuration{
		ID: id,
	}
	expectedStatusCode := http.StatusOK
	s.peripheralsMock.EXPECT().ByID(mock.Anything, id).Return(expectedConfiguration, nil).Once()
	client := http.Client{
		Timeout: time.Duration(time.Second),
	}
	apiPath := "http://localhost:8080/api/v1/configuration/" + id.String()

	// when
	resp, err := client.Get(apiPath)

	// then
	assertion.Nil(err)
	assertion.Equal(expectedStatusCode, resp.StatusCode)
}

func (s Suite) TestShouldGetConfiguration() {
	assertion := assert.New(s.T())

	// given
	id := "eda1a0ca-12a3-4d3e-97e6-5baa0c5a1b93"
	data := bytes.NewBuffer([]byte(fmt.Sprintf(`{"id":"%s"}`, id)))
	expectedConfiguration := storage.Configuration{
		ID: uuid.MustParse(id),
	}
	expectedStatusCode := http.StatusCreated
	s.peripheralsMock.EXPECT().Add(mock.Anything, expectedConfiguration).Return(nil).Once()
	client := http.Client{
		Timeout: time.Duration(time.Second),
	}
	apiPath := "http://localhost:8080/api/v1/configuration"

	// when
	resp, err := client.Post(apiPath, "application/json", data)

	// then
	assertion.Nil(err)
	assertion.Equal(expectedStatusCode, resp.StatusCode)
}

func (s Suite) TestShouldDeleteConfiguration() {
	assertion := assert.New(s.T())

	// given
	id := uuid.MustParse("eda1a0ca-12a3-4d3e-97e6-5baa0c5a1b93")
	expectedStatusCode := http.StatusNoContent
	s.peripheralsMock.EXPECT().DeleteOne(mock.Anything, id).Return(nil).Once()
	client := http.Client{
		Timeout: time.Duration(time.Second),
	}
	apiPath := "http://localhost:8080/api/v1/configuration/" + id.String()
	req, err := http.NewRequest(http.MethodDelete, apiPath, nil)
	assertion.Nil(err)

	// when
	resp, err := client.Do(req)

	// then
	assertion.Nil(err)
	assertion.Equal(expectedStatusCode, resp.StatusCode)
}

func (s Suite) TestShouldVerifyConfiguration() {
	assertion := assert.New(s.T())

	// given
	id := uuid.MustParse("eda1a0ca-12a3-4d3e-97e6-5baa0c5a1b93")
	expectedStatusCode := http.StatusOK
	s.peripheralsMock.EXPECT().Verify(mock.Anything, id).Return(true, nil).Once()
	client := http.Client{
		Timeout: time.Duration(time.Second),
	}
	apiPath := "http://localhost:8080/api/v1/configuration/" + id.String() + "/verify"

	// when
	resp, err := client.Post(apiPath, "application/json", nil)

	// then
	assertion.Nil(err)
	var actualResponse server.VerifyResponse
	err = json.NewDecoder(resp.Body).Decode(&actualResponse)
	assertion.Nil(err)
	assertion.True(actualResponse.Success)
	assertion.Equal(expectedStatusCode, resp.StatusCode)
}

func (s Suite) TestShouldReloadConfiguration() {
	assertion := assert.New(s.T())

	// given
	expectedStatusCode := http.StatusNoContent
	s.peripheralsMock.EXPECT().Reload(mock.Anything).Return(nil).Once()
	client := http.Client{
		Timeout: time.Duration(time.Second),
	}
	apiPath := "http://localhost:8080/api/v1/configurations/reload"

	// when
	resp, err := client.Post(apiPath, "application/json", nil)

	// then
	assertion.Nil(err)
	assertion.Equal(expectedStatusCode, resp.StatusCode)
}
