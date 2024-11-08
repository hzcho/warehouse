package service

import (
	"api_service/internal/domain/net/request"
	"context"
	"net/http"
)

type Product interface {
	GetAll(ctx context.Context, token string, filter request.GetAllFilter) (*http.Response, error)
	GetById(ctx context.Context, token string, id string) (*http.Response, error)
	Create(ctx context.Context, token string, req request.CreateProduct) (*http.Response, error)
	Update(ctx context.Context, token string, req request.UpdateUser) (*http.Response, error)
	UpdateCount(ctx context.Context, token string, req request.UpdateStockLevel) (*http.Response, error)
	Delete(ctx context.Context, token string, id string) (*http.Response, error)
}
