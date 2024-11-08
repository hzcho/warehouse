package usecase

import (
	"audit/internal/domain/model"
	"audit/internal/domain/net/request"
	"audit/internal/domain/net/response"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Operation interface {
	GetById(ctx context.Context, id primitive.ObjectID) (model.OperationLog, error)
	GetAll(ctx context.Context, filter request.GetAllFilter) (response.GetAllResponse, error)
	Save(ctx context.Context, log model.OperationLog) (primitive.ObjectID, error)
}
