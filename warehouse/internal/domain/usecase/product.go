package usecase

import (
	"context"
	"warehouse/internal/domain/model"
	"warehouse/internal/domain/net/request"
	"warehouse/internal/domain/net/response"
	"warehouse/pkg/token"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product interface {
	GetAll(ctx context.Context, claims token.AuthInfo, filter request.GetAllFilter) (response.Products, error)
	GetById(ctx context.Context, claims token.AuthInfo, id primitive.ObjectID) (model.Product, error)
	Create(ctx context.Context, claims token.AuthInfo, req request.CreateProduct) (primitive.ObjectID, error)
	Update(ctx context.Context, claims token.AuthInfo, req request.UpdateProduct) (model.Product, error)
	UpdateCount(ctx context.Context, claims token.AuthInfo, req request.UpdateStockLevel) (model.Product, error)
	Delete(ctx context.Context, claims token.AuthInfo, id primitive.ObjectID) (primitive.ObjectID, error)
}
