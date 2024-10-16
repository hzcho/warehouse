package dao

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `db:"id"`
	Login        string     `db:"login"`
	PassHash     string     `db:"pass_hash"`
	RoleId       uuid.UUID  `db:"role_id"`
	RefreshToken *string    `db:"refresh_token"`
	TokenExpiry  *time.Time `db:"token_expire"`
}

var UserColumns = []string{
	"id",
	"login",
	"pass_hash",
	"role_id",
	"refresh_token",
	"token_expiry",
}

type Role struct {
	ID   uuid.UUID `db:"id"`
	Role string    `db:"role_name"`
}

var RoleColumns = []string{
	"id",
	"role_name",
}
