package usecase

import (
	"context"
	"warehouse/pkg/token"
)

type Auth interface {
	VerifyToken(ctx context.Context, tkn string) (*token.AuthInfo, error)
}
