package handler

import "audit/internal/usecase"

type Handlers struct {
	*Operation
}

func NewHandlers(usecases *usecase.Usecases) *Handlers {
	return &Handlers{
		Operation: NewOperations(usecases),
	}
}
