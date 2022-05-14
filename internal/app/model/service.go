package model

import (
	"github.com/gorilla/mux"
	"net/http"
	"path"
)

type Service struct {
	Version        string            `yaml:"version"`
	Name           string            `yaml:"name"`
	EndpointPrefix string            `yaml:"prefix"`
	Port           int               `yaml:"port"`
	Hostname       string            `yaml:"hostname"`
	Routes         []Route           `yaml:"routes"`
	Headers        map[string]string `yaml:"headers"`
}

func CreateHandler(response Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(response.StatusCode)
		w.Write([]byte(response.Body))
	}
}

func (s Service) CreateRouter() *mux.Router {
	srv := mux.NewRouter()
	for _, route := range s.Routes {
		pattern := path.Join(s.EndpointPrefix, route.Endpoint)
		srv.HandleFunc(pattern, CreateHandler(route.Responses[0])).Methods(route.Method)
	}
	return srv
}
