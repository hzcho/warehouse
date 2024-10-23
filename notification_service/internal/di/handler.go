package di

import (
	"notification/internal/handler"
	"notification/internal/listener"
)

func NewHandlers(usecases *Usecases) *listener.Handlers {
	return &listener.Handlers{
		Warehouse: handler.NewWarehouse(usecases.Warehouse),
	}
}
