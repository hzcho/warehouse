package repository

import (
	"audit/internal/config"
	"audit/internal/domain/repository"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	repository.Operation
}

func NewRepositories(cfg *config.Config, db *mongo.Database, log *logrus.Logger) *Repositories {
	return &Repositories{
		Operation: NewOperation(db),
	}
}
