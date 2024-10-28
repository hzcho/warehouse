package usecase

import (
	"auth/internal/domain/net/request"
	"auth/internal/domain/net/response"
	"context"
)

type User interface {
	GetUsers(ctx context.Context, filter request.GetUsersFilter) (response.Users, error)
}
