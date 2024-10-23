package repository

import (
	"warehouse/internal/config"
	"warehouse/internal/domain/repository"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	repository.Product
	repository.FileStorage
}

func NewRepositories(cfg *config.Config, db *mongo.Database, log *logrus.Logger) *Repositories {
	return &Repositories{
		Product:     NewProduct(db, log),
		FileStorage: NewFileStorage(cfg.UploadDir),
	}
}
