package repository

import (
	"context"
	"warehouse/internal/domain/model"
	"warehouse/internal/domain/net/request"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category interface {
	GetAll(ctx context.Context, req request.GetCategories) ([]model.Category, error)
	GetByName(ctx context.Context, name string) (model.Category, error)
	Create(ctx context.Context, category model.Category) (primitive.ObjectID, error)
	Delete(ctx context.Context, id primitive.ObjectID) (primitive.ObjectID, error)
}
