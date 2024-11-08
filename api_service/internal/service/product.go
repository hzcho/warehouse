package service

import (
	"api_service/internal/domain/net/request"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	urlUtils "net/url"

	"github.com/sirupsen/logrus"
)

const (
	getAllURL        = "/products"
	createURL        = "/products"
	getByIdURL       = "/products"
	updateURL        = "/products"
	updateStockLevel = "/products/count"
	deleteURL        = "/products"
)

type Product struct {
	baseURL string
	client  *http.Client
	log     *logrus.Logger
}

func NewProduct(baseURL string, client *http.Client, log *logrus.Logger) *Product {
	return &Product{
		baseURL: baseURL,
		client:  client,
		log:     log,
	}
}

func (s *Product) GetAll(ctx context.Context, token string, filter request.GetAllFilter) (*http.Response, error) {
	url := s.baseURL + getAllURL

	values := urlUtils.Values{}

	if filter.ProductName != nil {
		values.Add("product_name", *filter.ProductName)
	}
	if filter.Page != nil {
		values.Add("page", fmt.Sprintf("%d", *filter.Page))
	}
	if filter.Limit != nil {
		values.Add("limit", fmt.Sprintf("%d", *filter.Limit))
	}

	if len(values) > 0 {
		url += "?" + values.Encode()
	}

	log := s.log.WithFields(logrus.Fields{
		"op":  "internal/usecase/warehouse/GetAll",
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

func (s *Product) Create(ctx context.Context, token string, req request.CreateProduct) (*http.Response, error) {
	url := s.baseURL + createURL

	log := s.log.WithFields(logrus.Fields{
		"op":  "internal/usecase/warehouse/Create",
		"url": url,
	})
	log.Info()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("name", req.Name)
	writer.WriteField("description", req.Description)
	writer.WriteField("category_name", req.CategoryName)
	writer.WriteField("price", fmt.Sprintf("%f", req.Price))
	writer.WriteField("stock_level", fmt.Sprintf("%d", req.StockLevel))
	writer.WriteField("min_stock_level", fmt.Sprintf("%d", req.MinStockLevel))
	writer.WriteField("manufacturer", req.Manufacturer)
	writer.WriteField("supplier", req.Supplier)
	writer.WriteField("weight", fmt.Sprintf("%f", req.Weight))
	writer.WriteField("height", fmt.Sprintf("%f", req.Dimensions.Height))
	writer.WriteField("length", fmt.Sprintf("%f", req.Dimensions.Length))
	writer.WriteField("width", fmt.Sprintf("%f", req.Dimensions.Width))

	for _, tag := range req.Tags {
		writer.WriteField("tags[]", tag)
	}

	for _, fileHeader := range req.Images {
		file, err := fileHeader.Open()
		if err != nil {
			log.Error(err)
			return nil, err
		}
		defer file.Close()

		part, err := writer.CreateFormFile("images", fileHeader.Filename)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		if _, err := io.Copy(part, file); err != nil {
			log.Error(err)
			return nil, err
		}
	}
	writer.Close()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	httpReq.Header.Set("Authorization", "Bearer "+token)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return resp, nil
}

func (s *Product) GetById(ctx context.Context, token string, id string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", s.baseURL+getByIdURL, id)

	log := s.log.WithFields(logrus.Fields{
		"op":  "internal/usecase/warehouse/GetById",
		"url": url,
	})
	log.Info()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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

func (s *Product) Update(ctx context.Context, token string, req request.UpdateUser) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", s.baseURL+getByIdURL, req.ID.Hex())

	log := s.log.WithFields(logrus.Fields{
		"op":  "internal/usecase/warehouse/UpdateUser",
		"url": url,
	})
	log.Info()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if req.Name != nil {
		writer.WriteField("name", *req.Name)
	}
	if req.Description != nil {
		writer.WriteField("description", *req.Description)
	}
	if req.CategoryName != nil {
		writer.WriteField("category_name", *req.CategoryName)
	}
	if req.Price != nil {
		writer.WriteField("price", fmt.Sprintf("%f", *req.Price))
	}
	if req.StockLevel != nil {
		writer.WriteField("stock_level", fmt.Sprintf("%d", *req.StockLevel))
	}
	if req.MinStockLevel != nil {
		writer.WriteField("min_stock_level", fmt.Sprintf("%d", *req.MinStockLevel))
	}
	if req.Manufacturer != nil {
		writer.WriteField("manufacturer", *req.Manufacturer)
	}
	if req.Supplier != nil {
		writer.WriteField("supplier", *req.Supplier)
	}
	if req.Weight != nil {
		writer.WriteField("weight", fmt.Sprintf("%f", *req.Weight))
	}
	if req.Dimensions != nil {
		writer.WriteField("height", fmt.Sprintf("%f", req.Dimensions.Height))
		writer.WriteField("length", fmt.Sprintf("%f", req.Dimensions.Length))
		writer.WriteField("width", fmt.Sprintf("%f", req.Dimensions.Width))
	}
	if req.Tags != nil {
		for _, tag := range *req.Tags {
			writer.WriteField("tags[]", tag)
		}
	}
	if req.IsActive != nil {
		writer.WriteField("is_active", fmt.Sprintf("%t", *req.IsActive))
	}

	if req.Images != nil {
		for _, fileHeader := range *req.Images {
			file, err := fileHeader.Open()
			if err != nil {
				log.Error(err)
				return nil, err
			}
			defer file.Close()

			part, err := writer.CreateFormFile("images", fileHeader.Filename)
			if err != nil {
				log.Error(err)
				return nil, err
			}
			if _, err := io.Copy(part, file); err != nil {
				log.Error(err)
				return nil, err
			}
		}
	}
	writer.Close()

	// Create and execute HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, url, body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	httpReq.Header.Set("Authorization", "Bearer "+token)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return resp, nil
}

func (s *Product) UpdateCount(ctx context.Context, token string, req request.UpdateStockLevel) (*http.Response, error) {
	url := s.baseURL + updateStockLevel + "/" + req.Id

	log := s.log.WithFields(logrus.Fields{
		"op":  "internal/usecase/warehouse/UpdateCount",
		"url": url,
	})
	log.Info()

	body, err := json.Marshal(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(body))
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

func (s *Product) Delete(ctx context.Context, token string, id string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", s.baseURL+deleteURL, id)

	log := s.log.WithFields(logrus.Fields{
		"op":  "internal/usecase/warehouse/Delete",
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
