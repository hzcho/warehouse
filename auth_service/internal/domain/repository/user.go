package repository

import (
	"auth/internal/domain/model"
	"auth/internal/domain/net/request"
	"context"

	"github.com/google/uuid"
)

type User interface {
	GetUsers(ctx context.Context, filter request.GetUsersFilter) ([]model.User, error)
	Get(ctx context.Context, login string) (model.User, error)
	Create(ctx context.Context, user model.User) (uuid.UUID, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
