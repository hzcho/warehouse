package handler

import "warehouse/internal/usecase"

type Handlers struct {
	*Product
	*Middleware
	*Category
}

func NewHandlers(usecases *usecase.Usecases) *Handlers {
	return &Handlers{
		Product:    NewProduct(usecases.Product),
		Middleware: NewMiddleware(usecases.Auth),
		Category:   NewCategory(usecases.Category),
	}
}
