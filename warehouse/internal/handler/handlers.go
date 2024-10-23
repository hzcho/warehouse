package handler

import "warehouse/internal/usecase"

type Handlers struct {
	*Product
}

func NewHandlers(usecases *usecase.Usecases) *Handlers {
	return &Handlers{
		Product: NewProduct(usecases.Product),
	}
}
