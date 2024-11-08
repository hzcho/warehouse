package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"time"
	"warehouse/internal/converter"
	"warehouse/internal/custom_errors.go"
	"warehouse/internal/domain/model"
	"warehouse/internal/domain/net/request"
	"warehouse/internal/domain/net/response"
	"warehouse/internal/domain/producer"
	"warehouse/internal/domain/repository"
	"warehouse/pkg/generate"
	"warehouse/pkg/token"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	minValueTopic = "min_value"
	saveOperation = "save_operation"
)

type Product struct {
	publisher    producer.Publisher
	productRepo  repository.Product
	categoryRepo repository.Category
	fileStorage  repository.FileStorage
	log          *logrus.Logger
	baseURL      string
}

func NewProduct(
	publisher producer.Publisher,
	productRepo repository.Product,
	categoryRepo repository.Category,
	fileStorage repository.FileStorage,
	log *logrus.Logger,
	baseURL string,
) *Product {
	return &Product{
		publisher:    publisher,
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		fileStorage:  fileStorage,
		log:          log,
		baseURL:      baseURL,
	}
}

func (u *Product) GetAll(ctx context.Context, claims token.AuthInfo, filter request.GetAllFilter) (response.Products, error) {
	log := u.log.WithFields(logrus.Fields{
		"op": "internal/usecase/product/GetAll",
	})

	if filter.Limit == nil || *filter.Limit <= 0 {
		defaultLimit := 10
		filter.Limit = &defaultLimit
	}

	if filter.Page == nil || *filter.Page < 0 {
		defaultPage := 0
		filter.Page = &defaultPage
	}

	if filter.ProductName != nil && *filter.ProductName == "" {
		filter.ProductName = nil
	}

	products, err := u.productRepo.GetAll(ctx, filter)
	if err != nil {
		log.Error(err)
		return response.Products{}, err
	}

	return response.Products{
		Page:     *filter.Page,
		Limit:    *filter.Limit,
		Products: products,
	}, nil
}

func (u *Product) GetById(ctx context.Context, claims token.AuthInfo, id primitive.ObjectID) (model.Product, error) {
	product, err := u.productRepo.GetById(ctx, id)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (u *Product) Create(ctx context.Context, claims token.AuthInfo, req request.CreateProduct) (primitive.ObjectID, error) {
	log := u.log.WithFields(logrus.Fields{
		"op":    "internal/usecase/product/Create",
		"login": claims.Login,
		"role":  claims.Role,
	})

	if claims.Role != "admin" && claims.Role != "manager" {
		err := errors.New("insufficient level of rights")
		log.Error(err)
		return primitive.NilObjectID, err
	}

	category, err := u.categoryRepo.GetByName(ctx, req.CategoryName)
	if err != nil {
		return primitive.NilObjectID, err
	}
	if category == (model.Category{}) {
		return primitive.NilObjectID, custom_errors.CategoryNotExist
	}

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

	operLog := model.OperationLog{
		UserId:        claims.UserID,
		ProductId:     id.Hex(),
		OperationType: "create",
		Timestamp:     time.Now(),
	}

	if err := u.publisher.Produce(saveOperation, operLog); err != nil {
		log.Error(err)
		return primitive.NilObjectID, err
	}

	return id, nil
}

func (u *Product) Update(ctx context.Context, claims token.AuthInfo, req request.UpdateProduct) (model.Product, error) {
	log := u.log.WithFields(logrus.Fields{
		"op": "internal/usecase/product/Update",
	})

	if claims.Role != "admin" && claims.Role != "manager" {
		err := errors.New("insufficient level of rights")
		log.Error(err)
		return model.Product{}, err
	}

	if req.CategoryName != nil {
		category, err := u.categoryRepo.GetByName(ctx, *req.CategoryName)
		if err != nil {
			return model.Product{}, err
		}
		if category == (model.Category{}) {
			return model.Product{}, custom_errors.CategoryNotExist
		}
	}

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
	if req.CategoryName != nil {
		updatedFields["category_name"] = *req.CategoryName
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
		UserId:        claims.UserID,
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

func (u *Product) UpdateCount(ctx context.Context, claims token.AuthInfo, req request.UpdateStockLevel) (model.Product, error) {
	log := u.log.WithFields(logrus.Fields{
		"op": "internal/usecase/product/UpdateCount",
	})
	product, err := u.productRepo.GetById(ctx, req.Id)
	if err != nil {
		log.Error(err)
		return model.Product{}, err
	}

	if product.StockLevel == nil || *product.StockLevel+req.StockLevel < 0 {
		err = errors.New("There is not enough quantity of goods")
		log.Error(err)
		return model.Product{}, err
	}

	*product.StockLevel += req.StockLevel

	updatedProduct, err := u.productRepo.Update(ctx, product)
	if err != nil {
		log.Error(err)
		return model.Product{}, err
	}

	operLog := model.OperationLog{
		UserId:        claims.UserID,
		ProductId:     updatedProduct.ID.Hex(),
		OperationType: "update",
		Timestamp:     time.Now(),
		UpdatedFields: map[string]interface{}{
			"stock_level": product,
		},
	}

	if err := u.publisher.Produce(saveOperation, operLog); err != nil {
		log.Error(err)
		return model.Product{}, err
	}

	return updatedProduct, nil
}

func (u *Product) Delete(ctx context.Context, claims token.AuthInfo, id primitive.ObjectID) (primitive.ObjectID, error) {
	log := u.log.WithFields(logrus.Fields{
		"op": "internal/usecase/product/Delete",
	})

	if claims.Role != "admin" && claims.Role != "manager" {
		err := errors.New("insufficient level of rights")
		log.Error(err)
		return primitive.NilObjectID, err
	}

	id, err := u.productRepo.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return primitive.NilObjectID, err
	}

	operLog := model.OperationLog{
		UserId:        claims.UserID,
		ProductId:     id.Hex(),
		OperationType: "delete",
		Timestamp:     time.Now(),
	}

	if err := u.publisher.Produce(saveOperation, operLog); err != nil {
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
