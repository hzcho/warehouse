package service

import (
	"api_service/internal/domain/net/request"
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	signinURL  = "/auth/signin"
	signupURL  = "/auth/signup"
	refreshURL = "/auth/refresh"
)

type Auth struct {
	baseURL string
	client  *http.Client
	log     *logrus.Logger
}

func NewAuth(baseURL string, client *http.Client, log *logrus.Logger) *Auth {
	return &Auth{
		baseURL: baseURL,
		client:  client,
		log:     log,
	}
}

func (s Auth) SignIn(ctx context.Context, req request.SignIn) (*http.Response, error) {
	url := s.baseURL + signinURL

	log := s.log.WithFields(logrus.Fields{
		"op":       "internal/usecase/auth/SignIn",
		"login":    req.Login,
		"password": req.Password,
		"url":      url,
	})
	log.Info()

	jsonData, err := json.Marshal(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(httpReq)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return resp, nil
}

func (s *Auth) SignUp(ctx context.Context, req request.SignUp) (*http.Response, error) {
	url := s.baseURL + signupURL

	log := s.log.WithFields(logrus.Fields{
		"op":       "internal/usecase/auth/SignUp",
		"login":    req.Login,
		"password": req.Password,
		"url":      url,
	})
	log.Info()

	jsonData, err := json.Marshal(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(httpReq)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return resp, nil
}

func (s *Auth) RefreshToken(ctx context.Context, req request.RefreshToken) (*http.Response, error) {
	url := s.baseURL + refreshURL

	log := s.log.WithFields(logrus.Fields{
		"op":  "internal/usecase/auth/RefreshToken",
		"url": url,
	})
	log.Info()

	jsonData, err := json.Marshal(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(httpReq)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return resp, nil
}
