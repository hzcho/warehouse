package usecase

import (
	"fmt"
	"warehouse/internal/config"
	"warehouse/internal/domain/producer"
	"warehouse/internal/domain/usecase"
	"warehouse/internal/repository"
	"warehouse/pkg/token"

	"github.com/sirupsen/logrus"
)

type Usecases struct {
	usecase.Product
	usecase.Auth
	usecase.Category
}

func NewUseCases(cfg *config.Config, publisher producer.Publisher, repos *repository.Repositories, log *logrus.Logger, tokenManager token.TokenManager) *Usecases {
	return &Usecases{
		Product: NewProduct(
			publisher,
			repos.Product,
			repos.Category,
			repos.FileStorage,
			log,
			fmt.Sprintf("http://%s:%s", cfg.Server.Host, cfg.Server.Port),
		),
		Auth:     NewAuth(log, cfg.Auth, tokenManager),
		Category: NewCategory(repos.Category, log),
	}
}
