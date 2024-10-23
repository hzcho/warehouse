package di

import (
	"net/http"
	servicesInterface "notification/internal/domain/service"
	"notification/internal/service"
)

type Services struct {
	servicesInterface.Auth
}

func NewServices(baseURL string, client *http.Client) *Services {
	return &Services{
		Auth: service.NewAuth(baseURL, client),
	}
}
