package app

import (
	"api_service/internal/config"
	"api_service/internal/handler"
	"api_service/internal/routing"
	"api_service/internal/server"
	"api_service/internal/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type App struct {
	server *server.Server
}

func NewApp(ctx context.Context, cfg *config.Config, log *logrus.Logger) *App {
	// tokenManager, err := token.NewManager(cfg.Auth.ATDuration, cfg.Auth.PrivateKeyPath, cfg.Auth.PublicKeyPath)
	// if err != nil {
	// 	panic(err)
	// }
	client := http.Client{}
	services := service.NewServices(cfg.URL, &client, log)
	handlers := handler.NewHandlers(services)
	router := gin.New()

	routing.InitRoutes(router, handlers)
	server := server.NewServer(&cfg.Server, router)

	return &App{
		server: server,
	}
}

func (a *App) Start() {
	a.server.Run()
}

func (a *App) Stop(ctx context.Context) {
	a.server.Stop(ctx)
}
