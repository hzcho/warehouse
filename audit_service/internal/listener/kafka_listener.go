package listener

import (
	"audit/internal/domain/consumer"
	topichandler "audit/internal/topic_handler"
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type TopicFunc func(context.Context, *kafka.Message) error

type KafkaListener struct {
	consumer      consumer.Consumer
	log           *logrus.Logger
	stop          chan struct{}
	topicHandlers map[string]TopicFunc
}

func New(consumer consumer.Consumer, log *logrus.Logger, handlers *topichandler.TopicHandlers) *KafkaListener {
	stop := make(chan struct{})
	topicHandlers := make(map[string]TopicFunc)

	topicHandlers["save_operation"] = handlers.Operation.Save

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
			return
		default:
			ms, err := k.consumer.Read()
			if err != nil {
				if err.Error() != "Local: Timed out" {
					k.log.Error(err)
				}
				continue
			}

			go func(ms *kafka.Message) {
				err := k.topicHandlers[*ms.TopicPartition.Topic](ctx, ms)
				if err != nil {
					k.log.Error(err)
				}
			}(ms)
		}
	}
}

func (k *KafkaListener) Stop() {
	k.stop <- struct{}{}
	defer close(k.stop)
}
