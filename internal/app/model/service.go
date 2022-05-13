package model

import (
	"log"
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

func (s Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	for _, route := range s.Routes {
		url := path.Join(s.EndpointPrefix, route.Endpoint)
		if route.Method == r.Method && url == r.URL.String() {
			w.Write([]byte(route.Responses[0].Body))
		}
	}
}
