package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

const (
	apiVer         = 1
	apiPathPattern = "/api/v%d/"
)

var (
	apiPath = fmt.Sprintf(apiPathPattern, apiVer)

	getConfigurationsPath   = path.Join(apiPath, "configurations")
	addConfigurationPath    = path.Join(apiPath, "configuration")
	getConfigurationPath    = path.Join(apiPath, "configuration/{id}")
	deleteConfigurationPath = path.Join(apiPath, "configuration/{id}")
	verifyConfigurationPath = path.Join(apiPath, "configuration/{id}/verify")
	reloadConfigurationPath = path.Join(apiPath, "configuration/reload")
	readDataSourcePath      = path.Join(apiPath, "data-source/{id}")
)

type ConfigurationService interface{}
type SchedulerService interface{}
type Conf interface {
	Addr() string
}

type Server struct {
	ctx context.Context

	confService  ConfigurationService
	schedService SchedulerService

	conf Conf
	log  *slog.Logger
}

func New(ctx context.Context, confService ConfigurationService, schedService SchedulerService, conf Conf, log *slog.Logger) Server {
	return Server{
		ctx:          ctx,
		confService:  confService,
		schedService: schedService,
		conf:         conf,
		log:          log,
	}
}

func (s Server) Start() error {
	m := mux.NewRouter()

	s.addEndpoint(m, addConfigurationPath, s.addConfiguration, http.MethodPost)
	s.addEndpoint(m, reloadConfigurationPath, s.reloadConfiguration, http.MethodPost)
	s.addEndpoint(m, getConfigurationPath, s.getConfiguration, http.MethodGet)
	s.addEndpoint(m, deleteConfigurationPath, s.deleteConfiguration, http.MethodDelete)
	s.addEndpoint(m, getConfigurationsPath, s.getConfigurations, http.MethodGet)
	s.addEndpoint(m, verifyConfigurationPath, s.verifyConfiguration, http.MethodGet)
	s.addEndpoint(m, readDataSourcePath, s.readDataSource, http.MethodGet)

	s.log.Info("starts listening on: " + s.conf.Addr())
	return http.ListenAndServe(s.conf.Addr(), m)
}

func (s Server) addEndpoint(mux *mux.Router, path string, handler func(http.ResponseWriter, *http.Request), methods ...string) {
	mux.HandleFunc(path, handler).Methods(methods...)

	for _, m := range methods {
		msg := fmt.Sprintf("registered handler for: %s %s", m, path)
		s.log.Info(msg)
	}
}

func (s Server) getConfigurations(w http.ResponseWriter, r *http.Request) {

}

func (s Server) addConfiguration(w http.ResponseWriter, r *http.Request) {

}

func (s Server) getConfiguration(w http.ResponseWriter, r *http.Request) {
}

func (s Server) deleteConfiguration(w http.ResponseWriter, r *http.Request) {

}

func (s Server) verifyConfiguration(w http.ResponseWriter, r *http.Request) {

}

func (s Server) reloadConfiguration(w http.ResponseWriter, r *http.Request) {

}

func (s Server) readDataSource(w http.ResponseWriter, r *http.Request) {

}
