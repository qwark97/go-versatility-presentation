package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/qwark97/go-versatility-presentation/hub/logger"
	"github.com/qwark97/go-versatility-presentation/hub/peripherals/storage"
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
	reloadConfigurationPath = path.Join(apiPath, "configurations/reload")
	readDataSourcePath      = path.Join(apiPath, "last-reading/{id}")
)

//go:generate mockery --name Peripherals --case underscore --with-expecter
type Peripherals interface {
	All(ctx context.Context) ([]storage.Configuration, error)
	Add(ctx context.Context, configuration storage.Configuration) error
	ByID(ctx context.Context, id uuid.UUID) (storage.Configuration, error)
	DeleteOne(ctx context.Context, id uuid.UUID) error
	Verify(ctx context.Context, id uuid.UUID) (bool, error)
	Reload(ctx context.Context) error
	LastReading(ctx context.Context, id uuid.UUID) (string, error)
}

//go:generate mockery --name Conf --case underscore --with-expecter
type Conf interface {
	Addr() string
	RequestTimeout() time.Duration
}

type Server struct {
	peripherals Peripherals

	conf Conf
	log  logger.Logger
}

func New(per Peripherals, conf Conf, log logger.Logger) Server {
	return Server{
		peripherals: per,
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
	s.addEndpoint(m, verifyConfigurationPath, s.verifyConfiguration, http.MethodPost)
	s.addEndpoint(m, readDataSourcePath, s.getLastReading, http.MethodGet)

	s.log.Info("starts listening on: %s", s.conf.Addr())
	return http.ListenAndServe(s.conf.Addr(), m)
}

func (s Server) addEndpoint(mux *mux.Router, path string, handler func(http.ResponseWriter, *http.Request), methods ...string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		s.log.Info("%s %s", r.Method, r.RequestURI)
		handler(w, r)
	}).Methods(methods...)

	for _, m := range methods {
		s.log.Info("registered handler for: %s %s", m, path)
	}
}

func (s Server) getConfigurations(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), s.conf.RequestTimeout())
	defer cancel()

	configurations, err := s.peripherals.All(ctx)
	if err != nil {
		s.log.Error("failed to get all configurations: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(configurations)
	if err != nil {
		s.log.Error("failed to send response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s Server) addConfiguration(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), s.conf.RequestTimeout())
	defer cancel()

	var configuration storage.Configuration
	err := json.NewDecoder(r.Body).Decode(&configuration)
	if err != nil {
		s.log.Error("failed to read request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.peripherals.Add(ctx, configuration)
	if err != nil {
		s.log.Error("failed to add new configration entity: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s Server) getConfiguration(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), s.conf.RequestTimeout())
	defer cancel()

	vars := mux.Vars(r)
	idVar := vars["id"]
	id, err := uuid.Parse(idVar)
	if err != nil {
		s.log.Error("invalid id: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	configuration, err := s.peripherals.ByID(ctx, id)
	if err != nil {
		s.log.Error("failed to get configuration: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if configuration.ID != id {
		s.log.Error("failed to find configuration by ID: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(configuration)
	if err != nil {
		s.log.Error("failed to send response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s Server) deleteConfiguration(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), s.conf.RequestTimeout())
	defer cancel()

	vars := mux.Vars(r)
	idVar := vars["id"]
	id, err := uuid.Parse(idVar)
	if err != nil {
		s.log.Error("invalid id: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.peripherals.DeleteOne(ctx, id)
	if err != nil {
		s.log.Error("failed to delete configuration: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s Server) verifyConfiguration(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), s.conf.RequestTimeout())
	defer cancel()

	vars := mux.Vars(r)
	idVar := vars["id"]
	id, err := uuid.Parse(idVar)
	if err != nil {
		s.log.Error("invalid id: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ok, err := s.peripherals.Verify(ctx, id)
	if err != nil {
		s.log.Error("failed to verify configuration: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := VerifyResponse{}
	response.Success = ok

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		s.log.Error("failed to send response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s Server) reloadConfiguration(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), s.conf.RequestTimeout())
	defer cancel()

	err := s.peripherals.Reload(ctx)
	if err != nil {
		s.log.Error("failed to reload configurations: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s Server) getLastReading(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), s.conf.RequestTimeout())
	defer cancel()

	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	idVar := vars["id"]
	id, err := uuid.Parse(idVar)
	if err != nil {
		s.log.Error("invalid id: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := s.peripherals.LastReading(ctx, id)
	if err != nil {
		s.log.Error("failed to get configuration: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(data))
	if err != nil {
		s.log.Error("failed to send response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
