package repository

import (
	"context"
	"fmt"
	"warehouse/internal/domain/model"
	"warehouse/internal/domain/net/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	categoriesCollection = "categories"
)

type Category struct {
	c *mongo.Collection
}

func NewCategory(db *mongo.Database) *Category {
	return &Category{
		c: db.Collection(categoriesCollection),
	}
}

func (r *Category) GetAll(ctx context.Context, req request.GetCategories) ([]model.Category, error) {
	var categories []model.Category
	mongoFilter := bson.M{}

	if req.Name != nil {
		mongoFilter["name"] = *req.Name
	}

	opts := options.Find()

	skip := int64(req.Page * req.Limit)
	limit := int64(req.Limit)
	opts.SetSkip(skip)
	opts.SetLimit(limit)

	cursor, err := r.c.Find(ctx, mongoFilter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var operLog model.Category
		if err := cursor.Decode(&operLog); err != nil {
			return nil, err
		}
		categories = append(categories, operLog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *Category) GetByName(ctx context.Context, name string) (model.Category, error) {
	filter := bson.D{{Key: "name", Value: name}}
	var category model.Category

	if err := r.c.FindOne(ctx, filter).Decode(&category); err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (r *Category) Create(ctx context.Context, category model.Category) (primitive.ObjectID, error) {
	result, err := r.c.InsertOne(ctx, category)
	if err != nil {
		return primitive.NilObjectID, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("не удалось получить ID вставленного документа")
	}

	return id, nil
}

func (r *Category) Delete(ctx context.Context, id primitive.ObjectID) (primitive.ObjectID, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := r.c.DeleteOne(ctx, filter)
	if err != nil {
		return primitive.NilObjectID, err
	}

	if result.DeletedCount == 0 {
		return primitive.NilObjectID, fmt.Errorf("документ с ID %s не найден", id.Hex())
	}

	return id, nil
}
