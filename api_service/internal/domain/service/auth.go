package service

import (
	"api_service/internal/domain/net/request"
	"context"
	"net/http"
)

type Auth interface {
	SignIn(ctx context.Context, req request.SignIn) (*http.Response, error)
	SignUp(ctx context.Context, req request.SignUp) (*http.Response, error)
	RefreshToken(ctx context.Context, req request.RefreshToken) (*http.Response, error)
}
