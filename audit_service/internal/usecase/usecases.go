package usecase

import (
	"audit/internal/domain/usecase"
	"audit/internal/repository"

	"github.com/sirupsen/logrus"
)

type Usecases struct {
	usecase.Operation
}

func NewUsecases(repos *repository.Repositories, log *logrus.Logger) *Usecases {
	return &Usecases{
		Operation: NewOperation(repos.Operation, log),
	}
}
