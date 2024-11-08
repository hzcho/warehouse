package usecase

import (
	"context"
	"warehouse/internal/domain/net/request"
	"warehouse/internal/domain/net/response"
	"warehouse/pkg/token"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category interface {
	GetAll(ctx context.Context, claims token.AuthInfo, req request.GetCategories) (response.GetCategories, error)
	Create(ctx context.Context, claims token.AuthInfo, req request.CreateCategory) (primitive.ObjectID, error)
	Delete(ctx context.Context, claims token.AuthInfo, id primitive.ObjectID) (primitive.ObjectID, error)
}
