package listener

import (
	"context"
	"notification/internal/domain/consumer"
	"notification/internal/domain/handler"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type TopicFunc func(context.Context, *kafka.Message) error

type Handlers struct {
	handler.Warehouse
}

type KafkaListener struct {
	consumer      consumer.Consumer
	log           *logrus.Logger
	stop          chan struct{}
	topicHandlers map[string]TopicFunc
}

func New(consumer consumer.Consumer, log *logrus.Logger, handlers *Handlers) *KafkaListener {
	stop := make(chan struct{})
	topicHandlers := make(map[string]TopicFunc)

	topicHandlers["min_value"] = handlers.Warehouse.MinValue

	return &KafkaListener{
		consumer:      consumer,
		log:           log,
		stop:          stop,
		topicHandlers: topicHandlers,
	}
}

func (k *KafkaListener) Start(ctx context.Context) {
	for {
		select {
		case <-k.stop:
			break
		default:
			ms, err := k.consumer.Read()
			if err != nil {
				if err.Error() != "Local: Timed out" {
					k.log.Error(err)
				}

				continue
			}

			err = k.topicHandlers[*ms.TopicPartition.Topic](ctx, ms)
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}

func (k *KafkaListener) Stop() {
	k.stop <- struct{}{}
	defer close(k.stop)
}
