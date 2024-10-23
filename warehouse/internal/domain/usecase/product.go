package usecase

import (
	"context"
	"warehouse/internal/domain/model"
	"warehouse/internal/domain/net/request"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product interface {
	Create(ctx context.Context, req request.CreateProduct) (primitive.ObjectID, error)
	GetById(ctx context.Context, id primitive.ObjectID) (model.Product, error)
	Update(ctx context.Context, req request.UpdateUser) (model.Product, error)
	Delete(ctx context.Context, id primitive.ObjectID) (primitive.ObjectID, error)
}
