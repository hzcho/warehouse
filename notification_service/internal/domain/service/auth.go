package service

import (
	"context"
	"notification/internal/domain/model"
)

type Auth interface {
	GetUsers(ctx context.Context, filter model.GetUserFilter) ([]model.User, error)
}
