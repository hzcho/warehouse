package handler

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Warehouse interface {
	MinValue(ctx context.Context, msg *kafka.Message) error
}
