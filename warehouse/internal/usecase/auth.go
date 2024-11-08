package usecase

import (
	"context"
	"warehouse/internal/config"
	"warehouse/pkg/token"

	"github.com/sirupsen/logrus"
)

type Auth struct {
	log          *logrus.Logger
	cfg          config.Auth
	tokenManager token.TokenManager
}

func NewAuth(log *logrus.Logger, cfg config.Auth, tokenManager token.TokenManager) *Auth {
	return &Auth{
		log:          log,
		cfg:          cfg,
		tokenManager: tokenManager,
	}
}

func (u *Auth) VerifyToken(ctx context.Context, tkn string) (*token.AuthInfo, error) {
	log := u.log.WithFields(logrus.Fields{
		"op": "internal/usecase/auth/VerifyToken",
	})

	claims, err := u.tokenManager.Parse(tkn)
	if err != nil {
		log.Errorf("access token: %v", err)
		return nil, err
	}

	return &claims, nil
}
