package handler

import "api_service/internal/service"

type Handlers struct {
	*Auth
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		Auth: NewAuth(services.Auth),
	}
}
