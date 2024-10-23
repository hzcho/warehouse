package repository

import (
	"context"
	"warehouse/internal/domain/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product interface {
	Save(ctx context.Context, product model.Product) (primitive.ObjectID, error)
	GetById(ctx context.Context, id primitive.ObjectID) (model.Product, error)
	Update(ctx context.Context, product model.Product) (model.Product, error)
	Delete(ctx context.Context, id primitive.ObjectID) (primitive.ObjectID, error)
}
