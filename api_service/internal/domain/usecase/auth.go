package usecase

import (
	"api_service/pkg/token"
	"context"
)

type Auth interface {
	VerifyToken(ctx context.Context, tkn string) (*token.AuthInfo, error)
}
