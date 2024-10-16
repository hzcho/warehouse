package controller

import "auth/internal/usecase"

type Controllers struct {
	*Auth
}

func NewControllers(usecases *usecase.Usecases) *Controllers {
	return &Controllers{
		Auth: NewAuth(usecases.Auth),
	}
}
