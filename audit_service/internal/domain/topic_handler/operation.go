package topichandler

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Operation interface {
	Save(ctx context.Context, msg *kafka.Message) error
}
