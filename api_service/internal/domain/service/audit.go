package service

import (
	"api_service/internal/domain/net/request"
	"context"
	"net/http"
)

type Audit interface {
	GetAll(ctx context.Context, token string, opt request.GetAllLogsFilter) (*http.Response, error)
}
