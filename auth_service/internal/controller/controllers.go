package controller

import "auth/internal/usecase"

type Controllers struct {
	*Auth
	*User
}

func NewControllers(usecases *usecase.Usecases) *Controllers {
	return &Controllers{
		Auth: NewAuth(usecases.Auth),
		User: NewUser(usecases.User),
	}
}
