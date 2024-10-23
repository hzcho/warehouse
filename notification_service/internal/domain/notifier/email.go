package notifier

import (
	"context"
	"notification/internal/domain/model"
)

type Email interface {
	SendMessage(ctx context.Context, message model.EmailMessage) error
}
