package app

import (
	"audit/internal/config"
	"audit/internal/consumer"
	"audit/internal/handler"
	"audit/internal/listener"
	"audit/internal/repository"
	"audit/internal/routing"
	"audit/internal/server"
	topichandler "audit/internal/topic_handler"
	"audit/internal/usecase"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	client   *mongo.Client
	server   *server.Server
	listener *listener.KafkaListener
}

func NewApp(ctx context.Context, cfg *config.Config, log *logrus.Logger) *App {
	path := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		cfg.Mongo.Username,
		cfg.Mongo.Password,
		cfg.Mongo.Host,
		cfg.Mongo.Port,
		cfg.Mongo.DBName)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(path))
	if err != nil {
		panic(err)
	}
	db := mongoClient.Database(cfg.Mongo.DBName)
	repos := repository.NewRepositories(cfg, db, log)
	usecases := usecase.NewUsecases(repos, log)
	handlers := handler.NewHandlers(usecases)
	router := gin.New()

	routing.InitRoutes(router, handlers)
	server := server.NewServer(&cfg.Server, router)

	//kafka
	consumer, err := consumer.NewKafkaConsumer(cfg.Consumer)
	if err != nil {
		panic(err)
	}
	topicHandlers := topichandler.NewTopicHandlers(usecases)

	listener := listener.New(consumer, log, topicHandlers)

	return &App{
		client:   mongoClient,
		server:   server,
		listener: listener,
	}
}

func (a *App) Start(ctx context.Context) {
	go a.server.Run()
	go a.listener.Start(ctx)
}

func (a *App) Stop(ctx context.Context) {
	a.server.Stop(ctx)
	a.client.Disconnect(ctx)
	a.listener.Stop()
}
