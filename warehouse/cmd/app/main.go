package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"warehouse/internal/app"
	"warehouse/internal/config"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	log := logrus.New()
	cfg, err := config.InitConfig("")
	if err != nil {
		panic(err)
	}

	app := app.NewApp(ctx, cfg, log)

	go func() {
		app.Start()
	}()
	log.Info("server is running")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	app.Stop(ctx)
	log.Info("server shutdown")
}
