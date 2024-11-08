package service

import (
	"api_service/internal/domain/net/request"
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category interface {
	GetAll(ctx context.Context, token string, req request.GetCategories) (*http.Response, error)
	Create(ctx context.Context, token string, req request.CreateCategory) (*http.Response, error)
	Delete(ctx context.Context, token string, id primitive.ObjectID) (*http.Response, error)
}
