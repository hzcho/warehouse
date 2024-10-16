package repository

import (
	"auth/internal/domain/model"
	"context"

	"github.com/google/uuid"
)

type User interface {
	Get(ctx context.Context, login string) (model.User, error)
	Create(ctx context.Context, user model.User) (uuid.UUID, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}