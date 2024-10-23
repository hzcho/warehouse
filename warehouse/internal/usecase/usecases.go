package usecase

import (
	"fmt"
	"warehouse/internal/config"
	"warehouse/internal/domain/usecase"
	"warehouse/internal/repository"

	"github.com/sirupsen/logrus"
)

type Usecases struct {
	usecase.Product
}

func NewUseCases(cfg *config.Config, repos *repository.Repositories, log *logrus.Logger) *Usecases {
	return &Usecases{
		Product: NewProduct(
			repos.Product,
			repos.FileStorage,
			log,
			fmt.Sprintf("http://%s:%s", cfg.Server.Host, cfg.Server.Port),
		),
	}
}
