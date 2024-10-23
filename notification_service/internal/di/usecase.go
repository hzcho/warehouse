package di

import (
	usecaseInterface "notification/internal/domain/usecase"
	"notification/internal/usecase"
)

type Usecases struct {
	usecaseInterface.Warehouse
}

func NewUsecases(notifiers *Notifiers, services *Services) *Usecases {
	return &Usecases{
		Warehouse: usecase.NewWarehouse(notifiers.Email, services.Auth),
	}
}
