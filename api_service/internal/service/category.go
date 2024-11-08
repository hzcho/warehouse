package service

import (
	"api_service/internal/domain/net/request"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	urlUtils "net/url"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	getCategoriesURL  = "/categories"
	createCategoryURL = "/categories"
	deleteCategoryURL = "/categories"
)

type Category struct {
	baseURL string
	client  *http.Client
	log     *logrus.Logger
}

func NewCategory(baseURL string, client *http.Client, log *logrus.Logger) *Category {
	return &Category{
		baseURL: baseURL,
		client:  client,
		log:     log,
	}
}

func (s *Category) GetAll(ctx context.Context, token string, req request.GetCategories) (*http.Response, error) {
	url := s.baseURL + getCategoriesURL

	values := urlUtils.Values{}

	if req.Name != nil {
		values.Add("name", *req.Name)
	}

	values.Add("page", fmt.Sprintf("%d", req.Page))
	values.Add("limit", fmt.Sprintf("%d", req.Limit))

	if len(values) > 0 {
		url += "?" + values.Encode()
	}

	log := s.log.WithFields(logrus.Fields{
		"op":  "internal/service/category/GetAll",
		"url": url,
	})
	log.Info()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+token)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return resp, nil
}

func (s *Category) Create(ctx context.Context, token string, req request.CreateCategory) (*http.Response, error) {
	url := s.baseURL + createCategoryURL

	log := s.log.WithFields(logrus.Fields{
		"op":  "internal/service/category/Category",
		"url": url,
	})
	log.Info()

	body, err := json.Marshal(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+token)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(httpReq)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return resp, nil
}

func (s *Category) Delete(ctx context.Context, token string, id primitive.ObjectID) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", s.baseURL+deleteCategoryURL, id.Hex())

	log := s.log.WithFields(logrus.Fields{
		"op":  "internal/service/category/Delete",
		"url": url,
	})
	log.Info()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+token)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return resp, nil
}
