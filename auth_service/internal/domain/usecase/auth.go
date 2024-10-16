package usecase

import (
	"auth/internal/domain/net/request"
	"auth/internal/domain/net/response"
	"context"
)

type Auth interface {
	SignUp(ctx context.Context, req request.SignUp) error
	SignIn(ctx context.Context, req request.SignIn) (response.Token, error)
	RefreshToken(ctx context.Context, token request.RefreshToken) (response.Token, error)
}
