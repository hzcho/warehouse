package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"notification/internal/domain/model"
)

const (
	getUsersURL = "/users"
)

type Auth struct {
	baseURL string
	client  *http.Client
}

func NewAuth(baseURL string, client *http.Client) *Auth {
	return &Auth{
		baseURL: baseURL,
		client:  client,
	}
}

func (s *Auth) GetUsers(ctx context.Context, filter model.GetUsersFilter) (model.Users, error) {
	url := s.baseURL + getUsersURL
	jsonData, err := json.Marshal(filter)
	if err != nil {
		return model.Users{}, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return model.Users{}, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return model.Users{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Users{}, fmt.Errorf("failed to get users: %s", resp.Status)
	}

	var users model.Users
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return model.Users{}, err
	}

	return users, nil
}
