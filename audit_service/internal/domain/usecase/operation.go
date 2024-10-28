package usecase

import (
	"audit/internal/domain/model"
	"audit/internal/domain/net/request"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Operation interface {
	GetById(ctx context.Context, id primitive.ObjectID) (model.OperationLog, error)
	GetAll(ctx context.Context, filter request.GetAllFilter) ([]model.OperationLog, error)
	Save(ctx context.Context, log model.OperationLog) (primitive.ObjectID, error)
}
