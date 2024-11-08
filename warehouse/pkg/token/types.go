package token

import (
	"github.com/golang-jwt/jwt/v5"
)

type AuthInfo struct {
	UserID string `json:"user_id"`
	Login  string `json:"login"`
	Role   string `json:"role"`
}

type tokenClaims struct {
	UserID string `json:"user_id"`
	Login  string `json:"login"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
