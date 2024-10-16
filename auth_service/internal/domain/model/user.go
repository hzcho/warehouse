package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Login        string
	Password     string
	Role         string
	RefreshToken *string
	TokenExpiry  *time.Time
}