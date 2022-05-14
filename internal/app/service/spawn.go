package service

import (
	"fmt"
	"log"
	"net/http"
	"travesty/internal/app/model"
)

type Spawner interface {
	Start(service model.Service)
}

type serviceSpawner struct {
}

func NewSpawner() Spawner {
	return serviceSpawner{}
}

func (s serviceSpawner) Start(service model.Service) {
	addr := fmt.Sprintf("%s:%d", service.Hostname, service.Port)
	log.Printf("Starting mock service [%s] at %s", service.Name, addr)
	go http.ListenAndServe(addr, service.CreateRouter())
}
