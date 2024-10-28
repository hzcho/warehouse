package handler

import "api_service/internal/service"

type Handlers struct {
	*Auth
	*Warehouse
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		Auth:      NewAuth(services.Auth),
		Warehouse: NewWarehouse(services.Warehouse),
	}
}
