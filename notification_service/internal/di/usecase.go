package di

import (
	usecaseInterface "notification/internal/domain/usecase"
	"notification/internal/usecase"

	"github.com/sirupsen/logrus"
)

type Usecases struct {
	usecaseInterface.Warehouse
}

func NewUsecases(notifiers *Notifiers, services *Services, log *logrus.Logger) *Usecases {
	return &Usecases{
		Warehouse: usecase.NewWarehouse(notifiers.Email, services.Auth, log),
	}
}
