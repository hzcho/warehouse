package usecase

import (
	"api_service/internal/config"
	"api_service/internal/domain/usecase"
	"api_service/pkg/token"

	"github.com/sirupsen/logrus"
)

type Usecases struct {
	usecase.Auth
}

func NewUsecases(log *logrus.Logger, cfg *config.Config, tokenManager token.TokenManager) *Usecases {
	return &Usecases{
		Auth: NewAuth(log, cfg.Auth, tokenManager),
	}
}
