package usecase

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"
	"warehouse/internal/converter"
	"warehouse/internal/domain/model"
	"warehouse/internal/domain/net/request"
	"warehouse/internal/domain/producer"
	"warehouse/internal/domain/repository"
	"warehouse/pkg/generate"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	minValueTopic = "min_value"
	saveOperation = "save_operation"
)

type Product struct {
	publisher   producer.Publisher
	productRepo repository.Product
	fileStorage repository.FileStorage
	log         *logrus.Logger
	baseURL     string
}

func NewProduct(publisher producer.Publisher, productRepo repository.Product, fileStorage repository.FileStorage, log *logrus.Logger, baseURL string) *Product {
	return &Product{
		publisher:   publisher,
		productRepo: productRepo,
		fileStorage: fileStorage,
		log:         log,
		baseURL:     baseURL,
	}
}

func (u *Product) GetById(ctx context.Context, id primitive.ObjectID) (model.Product, error) {
	product, err := u.productRepo.GetById(ctx, id)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (u *Product) Create(ctx context.Context, req request.CreateProduct) (primitive.ObjectID, error) {
	log := u.log.WithFields(logrus.Fields{
		"op": "internal/usecase/product/Create",
	})

	imageURLs, err := u.saveImages(ctx, req.Images)
	if err != nil {
		log.Error(err)
		return primitive.NilObjectID, err
	}

	product := converter.ProductFromCreate(req)
	newID := primitive.NewObjectID()
	product.ID = &newID
	now := time.Now()
	product.CreatedAt = &now
	product.UpdatedAt = &now
	product.ImageURLs = &imageURLs
	active := true
	product.IsActive = &active

	id, err := u.productRepo.Save(ctx, product)
	if err != nil {
		log.Error(err)
		return primitive.NilObjectID, err
	}

	return id, nil
}

func (u *Product) Update(ctx context.Context, req request.UpdateUser) (model.Product, error) {
	log := u.log.WithFields(logrus.Fields{
		"op": "internal/usecase/product/Update",
	})

	product := converter.ProductFromUpdate(req)
	now := time.Now()
	product.UpdatedAt = &now

	updatedProduct, err := u.productRepo.Update(ctx, product)
	if err != nil {
		log.Error(err)
		return model.Product{}, err
	}

	if *updatedProduct.StockLevel <= *updatedProduct.MinStockLevel {
		minValue := model.MinValue{
			ProductName:   *updatedProduct.Name,
			StockLevel:    *updatedProduct.StockLevel,
			MinStockLevel: *updatedProduct.MinStockLevel,
		}
		if err := u.publisher.Produce(minValueTopic, minValue); err != nil {
			log.Error(err)
			return model.Product{}, err
		}
	}

	updatedFields := make(map[string]interface{})

	if req.Name != nil {
		updatedFields["name"] = *req.Name
	}
	if req.Description != nil {
		updatedFields["description"] = *req.Description
	}
	if req.CategoryId != nil {
		updatedFields["category_id"] = *req.CategoryId
	}
	if req.Price != nil {
		updatedFields["price"] = *req.Price
	}
	if req.StockLevel != nil {
		updatedFields["stock_level"] = *req.StockLevel
	}
	if req.MinStockLevel != nil {
		updatedFields["min_stock_level"] = *req.MinStockLevel
	}
	if req.Manufacturer != nil {
		updatedFields["manufacturer"] = *req.Manufacturer
	}
	if req.Supplier != nil {
		updatedFields["supplier"] = *req.Supplier
	}
	if req.Weight != nil {
		updatedFields["weight"] = *req.Weight
	}
	if req.Dimensions != nil {
		updatedFields["dimensions"] = *req.Dimensions
	}
	if req.Tags != nil {
		updatedFields["tags"] = *req.Tags
	}
	if req.ImageURLs != nil {
		updatedFields["image_urls"] = *req.ImageURLs
	}
	if req.IsActive != nil {
		updatedFields["is_active"] = *req.IsActive
	}

	operLog := model.OperationLog{
		UserId:        "popa_enota",
		ProductId:     updatedProduct.ID.Hex(),
		OperationType: "update",
		Timestamp:     time.Now(),
		UpdatedFields: updatedFields,
	}

	if err := u.publisher.Produce(saveOperation, operLog); err != nil {
		log.Error(err)
		return model.Product{}, err
	}

	return updatedProduct, nil
}

func (u *Product) Delete(ctx context.Context, id primitive.ObjectID) (primitive.ObjectID, error) {
	log := u.log.WithFields(logrus.Fields{
		"op": "internal/usecase/product/Delete",
	})
	id, err := u.productRepo.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return primitive.NilObjectID, err
	}

	return id, nil
}

func (u *Product) formImageURL(relativePath string) (string, error) {
	imgURL, err := url.JoinPath(u.baseURL, relativePath)
	if err != nil {
		return "", fmt.Errorf("failed to create image URL: %w", err)
	}
	return imgURL, nil
}

func (u *Product) saveImages(ctx context.Context, images []io.Reader) ([]string, error) {
	var imageURLs []string
	for _, im := range images {
		baseName, err := generate.RandomString(20)
		if err != nil {
			return nil, err
		}
		name := baseName

		for i := 0; ; i++ {
			isExists, err := u.fileStorage.IsExists(name)
			if err != nil {
				return nil, err
			}
			if !isExists {
				break
			}

			name = fmt.Sprintf("%s_%d", baseName, i)
		}

		if err = u.fileStorage.SaveFile(ctx, name, im); err != nil {
			return nil, err
		}

		url, err := url.JoinPath(u.baseURL, u.fileStorage.AbsPath(), name)
		if err != nil {
			return nil, err
		}
		imageURLs = append(imageURLs, url)
	}

	return imageURLs, nil
}
