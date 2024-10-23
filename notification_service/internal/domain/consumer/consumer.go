package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer interface {
	Read() (*kafka.Message, error)
}
