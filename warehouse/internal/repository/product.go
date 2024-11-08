package repository

import (
	"context"
	"fmt"
	"warehouse/internal/domain/model"
	"warehouse/internal/domain/net/request"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

const (
	productsCollection = "products"
)

type Product struct {
	c   *mongo.Collection
	log *logrus.Logger
}

func NewProduct(db *mongo.Database, log *logrus.Logger) *Product {
	return &Product{
		c:   db.Collection(productsCollection),
		log: log,
	}
}

func (r *Product) GetAll(ctx context.Context, filter request.GetAllFilter) ([]model.Product, error) {
	var products []model.Product
	mongoFilter := bson.M{}

	if filter.ProductName != nil {
		mongoFilter["name"] = *filter.ProductName
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
		var operLog model.Product
		if err := cursor.Decode(&operLog); err != nil {
			return nil, err
		}
		products = append(products, operLog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *Product) GetById(ctx context.Context, id primitive.ObjectID) (model.Product, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	var product model.Product

	if err := r.c.FindOne(ctx, filter).Decode(&product); err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (r *Product) Save(ctx context.Context, product model.Product) (primitive.ObjectID, error) {
	result, err := r.c.InsertOne(ctx, product)
	if err != nil {
		return primitive.NilObjectID, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("не удалось получить ID вставленного документа")
	}

	return id, nil
}

func (r *Product) Update(ctx context.Context, product model.Product) (model.Product, error) {
	if product.ID == nil {
		return model.Product{}, fmt.Errorf("ID продукта не указан")
	}

	update := bson.D{}

	if product.Name != nil {
		update = append(update, bson.E{Key: "name", Value: product.Name})
	}
	if product.Description != nil {
		update = append(update, bson.E{Key: "description", Value: product.Description})
	}
	if product.CategoryName != nil {
		update = append(update, bson.E{Key: "category_name", Value: product.CategoryName})
	}
	if product.Price != nil {
		update = append(update, bson.E{Key: "price", Value: product.Price})
	}
	if product.StockLevel != nil {
		update = append(update, bson.E{Key: "stock_level", Value: product.StockLevel})
	}
	if product.MinStockLevel != nil {
		update = append(update, bson.E{Key: "min_stock_level", Value: product.MinStockLevel})
	}
	if product.Manufacturer != nil {
		update = append(update, bson.E{Key: "manufacturer", Value: product.Manufacturer})
	}
	if product.Supplier != nil {
		update = append(update, bson.E{Key: "supplier", Value: product.Supplier})
	}
	if product.Weight != nil {
		update = append(update, bson.E{Key: "weight", Value: product.Weight})
	}
	if product.Dimensions != nil {
		update = append(update, bson.E{Key: "dimensions", Value: product.Dimensions})
	}
	if product.ImageURLs != nil {
		update = append(update, bson.E{Key: "images", Value: product.ImageURLs})
	}
	if product.Tags != nil {
		update = append(update, bson.E{Key: "tags", Value: product.Tags})
	}
	if product.IsActive != nil {
		update = append(update, bson.E{Key: "is_active", Value: product.IsActive})
	}

	if len(update) == 0 {
		return model.Product{}, fmt.Errorf("нет полей для обновления")
	}

	filter := bson.D{{Key: "_id", Value: product.ID}}

	var updatedProduct model.Product
	err := r.c.FindOneAndUpdate(ctx, filter, bson.D{{Key: "$set", Value: update}}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedProduct)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Product{}, fmt.Errorf("продукт с ID %s не найден", product.ID.Hex())
		}
		return model.Product{}, err
	}

	return updatedProduct, nil
}

func (r *Product) Delete(ctx context.Context, id primitive.ObjectID) (primitive.ObjectID, error) {
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
