package repository

import (
	"auth/internal/domain/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	repository.User
}

func NewRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		User: NewUser(pool),
	}
}
