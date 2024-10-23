package di

import (
	"notification/internal/config"
	"notification/internal/consumer"
	consumerInterface "notification/internal/domain/consumer"
)

type Consumers struct {
	consumerInterface.Consumer
}

func NewConsumers(cfg *config.Config) (*Consumers, error) {
	kafkaConsumer, err := consumer.NewKafkaConsumer(cfg.Consumer)
	if err != nil {
		return nil, err
	}

	return &Consumers{
		Consumer: kafkaConsumer,
	}, nil
}
