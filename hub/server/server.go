package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"path"
	"time"

	"github.com/gorilla/mux"
	"github.com/qwark97/go-versatility-presentation/hub/peripherals/model"
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

type Peripherals interface {
	All(ctx context.Context) ([]model.Configuration, error)
	Add(ctx context.Context, configuration model.Configuration) error
}
type Scheduler interface{}
type Conf interface {
	Addr() string
	RequestTimeout() time.Duration
}

type Server struct {
	peripherals Peripherals
	scheduler   Scheduler

	conf Conf
	log  *slog.Logger
}

func New(peripherals Peripherals, scheduler Scheduler, conf Conf, log *slog.Logger) Server {
	return Server{
		peripherals: peripherals,
		scheduler:   scheduler,
		conf:        conf,
		log:         log,
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
	s.addEndpoint(m, readDataSourcePath, s.readData, http.MethodGet)

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
	w.Header().Add("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(r.Context(), s.conf.RequestTimeout())
	defer cancel()

	configurations, err := s.peripherals.All(ctx)
	if err != nil {
		s.log.Error(fmt.Sprintf("failed to get all configurations: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(configurations)
	if err != nil {
		s.log.Error(fmt.Sprintf("failed to send response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s Server) addConfiguration(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), s.conf.RequestTimeout())
	defer cancel()

	var configuration model.Configuration
	err := json.NewDecoder(r.Body).Decode(&configuration)
	if err != nil {
		s.log.Error(fmt.Sprintf("failed to read request body: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.peripherals.Add(ctx, configuration)
	if err != nil {
		s.log.Error(fmt.Sprintf("failed to add new configration entity: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s Server) getConfiguration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
}

func (s Server) deleteConfiguration(w http.ResponseWriter, r *http.Request) {

}

func (s Server) verifyConfiguration(w http.ResponseWriter, r *http.Request) {

}

func (s Server) reloadConfiguration(w http.ResponseWriter, r *http.Request) {

}

func (s Server) readData(w http.ResponseWriter, r *http.Request) {

}
