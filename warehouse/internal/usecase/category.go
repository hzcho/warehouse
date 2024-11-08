package usecase

import (
	"context"
	"errors"
	"warehouse/internal/domain/model"
	"warehouse/internal/domain/net/request"
	"warehouse/internal/domain/net/response"
	"warehouse/internal/domain/repository"
	"warehouse/pkg/token"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	categoryRepo repository.Category
	log          *logrus.Logger
}

func NewCategory(categoryRepo repository.Category, log *logrus.Logger) *Category {
	return &Category{
		categoryRepo: categoryRepo,
		log:          log,
	}
}

func (u *Category) GetAll(ctx context.Context, claims token.AuthInfo, req request.GetCategories) (response.GetCategories, error) {
	categories, err := u.categoryRepo.GetAll(ctx, req)
	if err != nil {
		return response.GetCategories{}, err
	}

	return response.GetCategories{
		Page:       req.Page,
		Limit:      req.Limit,
		Categories: categories,
	}, nil
}

func (u *Category) Create(ctx context.Context, claims token.AuthInfo, req request.CreateCategory) (primitive.ObjectID, error) {
	log := u.log.WithFields(logrus.Fields{
		"op": "internal/usecase/category/Create",
	})

	if claims.Role != "admin" && claims.Role != "manager" {
		err := errors.New("insufficient level of rights")
		log.Error(err)
		return primitive.NilObjectID, err
	}

	category := model.Category{
		ID:          primitive.NewObjectID(),
		Name:        req.Name,
		Description: req.Description,
	}

	id, err := u.categoryRepo.Create(ctx, category)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return id, nil
}

func (u *Category) Delete(ctx context.Context, claims token.AuthInfo, id primitive.ObjectID) (primitive.ObjectID, error) {
	log := u.log.WithFields(logrus.Fields{
		"op": "internal/usecase/category/Delete",
	})

	if claims.Role != "admin" && claims.Role != "manager" {
		err := errors.New("insufficient level of rights")
		log.Error(err)
		return primitive.NilObjectID, err
	}

	id, err := u.categoryRepo.Delete(ctx, id)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return id, nil
}
