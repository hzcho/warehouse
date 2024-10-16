package usecase

import (
	"auth/internal/config"
	"auth/internal/domain/usecase"
	"auth/internal/repository"
	"auth/pkg/token"

	"github.com/sirupsen/logrus"
)

type Usecases struct {
	usecase.Auth
}

func NewUsecases(repos *repository.Repositories, log *logrus.Logger, cfg config.Auth, tokenManager token.TokenManager) *Usecases {
	return &Usecases{
		Auth: NewAuth(repos.User, log, cfg, tokenManager),
	}
}
