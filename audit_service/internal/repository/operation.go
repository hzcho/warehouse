package repository

import (
	"audit/internal/domain/model"
	"audit/internal/domain/net/request"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	operationCollection = "operation"
)

type Operation struct {
	c *mongo.Collection
}

func NewOperation(db *mongo.Database) *Operation {
	return &Operation{
		c: db.Collection(operationCollection),
	}
}

func (r *Operation) GetById(ctx context.Context, id primitive.ObjectID) (model.OperationLog, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	var product model.OperationLog

	if err := r.c.FindOne(ctx, filter).Decode(&product); err != nil {
		return model.OperationLog{}, err
	}

	return product, nil
}

func (r *Operation) GetAll(ctx context.Context, filter request.GetAllFilter) ([]model.OperationLog, error) {
	var operLogs []model.OperationLog

	mongoFilter := bson.M{}
	if filter.UserId != nil {
		mongoFilter["user_id"] = *filter.UserId
	}
	if filter.ProductId != nil {
		mongoFilter["product_id"] = *filter.ProductId
	}
	if filter.OperationType != nil {
		mongoFilter["operation_type"] = *filter.OperationType
	}
	if filter.Timestamp != nil {
		mongoFilter["timestamp"] = *filter.Timestamp
	}

	opts := options.Find()
	if filter.Page != nil && filter.Limit != nil {
		skip := int64(*filter.Page * *filter.Limit)
		limit := int64(*filter.Limit)
		opts.SetSkip(skip)
		opts.SetLimit(limit)
	}

	cursor, err := r.c.Find(ctx, mongoFilter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var operLog model.OperationLog
		if err := cursor.Decode(&operLog); err != nil {
			return nil, err
		}
		operLogs = append(operLogs, operLog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return operLogs, nil
}

func (r *Operation) Save(ctx context.Context, log model.OperationLog) (primitive.ObjectID, error) {
	result, err := r.c.InsertOne(ctx, log)
	if err != nil {
		return primitive.NilObjectID, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("не удалось получить ID вставленного документа")
	}

	return id, nil
}
