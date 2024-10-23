package app

import (
	"context"
	"fmt"
	"warehouse/internal/config"
	"warehouse/internal/handler"
	"warehouse/internal/repository"
	"warehouse/internal/routing"
	"warehouse/internal/server"
	"warehouse/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	client *mongo.Client
	server *server.Server
}

func NewApp(ctx context.Context, cfg *config.Config, log *logrus.Logger) *App {
	path := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		cfg.Mongo.Username,
		cfg.Mongo.Password,
		cfg.Mongo.Host,
		cfg.Mongo.Port,
		cfg.Mongo.DBName)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(path))
	if err != nil {
		panic(err)
	}
	db := client.Database(cfg.Mongo.DBName)
	repos := repository.NewRepositories(cfg, db, log)
	usecases := usecase.NewUseCases(cfg, repos, log)
	handlers := handler.NewHandlers(usecases)
	router := gin.New()

	routing.InitRoutes(router, handlers)
	server := server.NewServer(&cfg.Server, router)

	return &App{
		client: client,
		server: server,
	}
}

func (a *App) Start() {
	a.server.Run()
}

func (a *App) Stop(ctx context.Context) {
	a.server.Stop(ctx)
	a.client.Disconnect(ctx)
}
