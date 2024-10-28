package service

import (
	"api_service/internal/domain/net/request"
	"context"
	"net/http"
)

type Warehouse interface {
	Create(ctx context.Context, req request.CreateProduct) (*http.Response, error)
	GetById(ctx context.Context, id string) (*http.Response, error)
	Update(ctx context.Context, req request.UpdateUser) (*http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
}
