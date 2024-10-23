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
	"warehouse/internal/domain/repository"
	"warehouse/pkg/generate"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	productRepo repository.Product
	fileStorage repository.FileStorage
	log         *logrus.Logger
	baseURL     string
}

func (u *Product) GetById(ctx context.Context, id primitive.ObjectID) (model.Product, error) {
	product, err := u.productRepo.GetById(ctx, id)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func NewProduct(productRepo repository.Product, fileStorage repository.FileStorage, log *logrus.Logger, baseURL string) *Product {
	return &Product{
		productRepo: productRepo,
		fileStorage: fileStorage,
		log:         log,
		baseURL:     baseURL,
	}
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
		"op": "internal/usecase/product/Delete",
	})

	product := converter.ProductFromUpdate(req)
	now := time.Now()
	product.UpdatedAt = &now

	updatedProduct, err := u.productRepo.Update(ctx, product)
	if err != nil {
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
