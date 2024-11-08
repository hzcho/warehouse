package service

import (
	"api_service/internal/domain/net/request"
	"context"
	"fmt"
	"net/http"
	urlUtils "net/url"

	"github.com/sirupsen/logrus"
)

const (
	getLogsURL = "/operations"
)

type Audit struct {
	baseURL string
	client  *http.Client
	log     *logrus.Logger
}

func NewAudit(baseURL string, client *http.Client, log *logrus.Logger) *Audit {
	return &Audit{
		baseURL: baseURL,
		client:  client,
		log:     log,
	}
}

func (s *Audit) GetAll(ctx context.Context, token string, opt request.GetAllLogsFilter) (*http.Response, error) {
	log := s.log.WithFields(logrus.Fields{
		"op": "internal/usecase/audit/GetAll",
	})

	url := s.baseURL + getLogsURL

	values := urlUtils.Values{}

	values.Add("page", fmt.Sprintf("%d", opt.Page))
	values.Add("limit", fmt.Sprintf("%d", opt.Limit))
	if opt.UserId != nil {
		values.Add("user_id", fmt.Sprintf("%d", opt.UserId))
	}
	if opt.ProductId != nil {
		values.Add("product_id", fmt.Sprintf("%d", opt.ProductId))
	}
	if opt.OperationType != nil {
		values.Add("operation_type", fmt.Sprintf("%d", opt.OperationType))
	}
	if opt.Timestamp != nil {
		values.Add("timestamp", fmt.Sprintf("%v", opt.Timestamp))
	}

	if len(values) > 0 {
		url += "?" + values.Encode()
	}

	log.Infof("url: %s", url)

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
