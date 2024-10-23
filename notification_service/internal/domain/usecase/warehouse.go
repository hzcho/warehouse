package usecase

import (
	"context"
	"notification/internal/domain/model"
)

type Warehouse interface {
	MinValue(ctx context.Context, event model.MinValue) error
}
