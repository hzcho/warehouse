package service

import (
	"context"
	"notification/internal/domain/model"
)

type Auth interface {
	GetUsers(ctx context.Context, filter model.GetUsersFilter) (model.Users, error)
}
