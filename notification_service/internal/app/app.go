package app

import (
	"context"
	"net/http"
	"notification/internal/config"
	"notification/internal/di"
	"notification/internal/listener"

	"github.com/sirupsen/logrus"
)

type App struct {
	listener *listener.KafkaListener
}

func NewApp(cfg *config.Config, log *logrus.Logger) *App {
	client := &http.Client{}

	listener, err := di.InitListener(cfg, log, client)
	if err != nil {
		panic(err)
	}

	return &App{
		listener: listener,
	}
}

func (a *App) Start(ctx context.Context) {
	a.listener.Start(ctx)
}

func (a *App) Stop() {
	a.listener.Stop()
}
