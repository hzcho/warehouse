package di

import (
	"net/http"
	"notification/internal/config"
	"notification/internal/listener"

	"github.com/sirupsen/logrus"
)

func InitListener(cfg *config.Config, log *logrus.Logger, client *http.Client) (*listener.KafkaListener, error) {
	notifiers := NewNotifiers(cfg)
	services := NewServices(cfg.URLs.Auth, client)
	usecases := NewUsecases(notifiers, services)
	handlers := NewHandlers(usecases)
	consumers, err := NewConsumers(cfg)
	if err != nil {
		return nil, err
	}
	return listener.New(consumers.Consumer, log, handlers), nil
}
