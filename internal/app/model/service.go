package model

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path"
	"text/template"
	"travesty/internal/app/util/functions"
)

type Service struct {
	Version        string            `yaml:"version"`
	Name           string            `yaml:"name"`
	EndpointPrefix string            `yaml:"prefix"`
	Port           int               `yaml:"port"`
	Hostname       string            `yaml:"hostname"`
	Routes         map[string]Route  `yaml:"routes"`
	Headers        map[string]string `yaml:"headers"`
}

func CreateHandler(response Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.New("response")
		tmpl.Funcs(functions.GlobalFunctions())
		w.WriteHeader(response.StatusCode)
		tmpl.Parse(response.Body)
		tmpl.Execute(w, nil)
	}
}

func (s Service) CreateRouter() *mux.Router {
	srv := mux.NewRouter()
	for _, route := range s.Routes {
		pattern := path.Join(s.EndpointPrefix, route.Endpoint)
		log.Println("adding route for", route.Method, pattern)
		srv.HandleFunc(pattern, CreateHandler(route.Responses[0])).Methods(route.Method)
	}
	return srv
}
