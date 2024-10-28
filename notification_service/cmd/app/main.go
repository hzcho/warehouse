package main

import (
	"context"
	"notification/internal/app"
	"notification/internal/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	log := logrus.New()
	cfg, err := config.InitConfig("")
	if err != nil {
		panic(err)
	}

	app := app.NewApp(cfg, log)

	go func() {
		app.Start(ctx)
	}()
	log.Info("server is running")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	app.Stop()
	log.Info("server shutdown")
}
