package handler

import (
	"api_service/internal/service"
	"api_service/internal/usecase"
)

type Handlers struct {
	*Auth
	*Product
	*Middleware
	*Audit
	*Category
}

func NewHandlers(services *service.Services, usecases *usecase.Usecases) *Handlers {
	return &Handlers{
		Auth:       NewAuth(services.Auth),
		Product:    NewProduct(services.Product),
		Middleware: NewMiddleware(usecases.Auth),
		Audit:      NewAudit(services.Audit),
		Category:   NewCategory(services.Category),
	}
}
