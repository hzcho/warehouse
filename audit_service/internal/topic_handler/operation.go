package topichandler

import (
	"audit/internal/domain/model"
	"audit/internal/domain/usecase"
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Operation struct {
	operUseCase usecase.Operation
}

func NewOperation(operUseCase usecase.Operation) *Operation {
	return &Operation{
		operUseCase: operUseCase,
	}
}

func (h *Operation) Save(ctx context.Context, msg *kafka.Message) error {
	var event model.OperationLog
	if err := json.Unmarshal(msg.Value, &event); err != nil {
		return err
	}

	if _, err := h.operUseCase.Save(ctx, event); err != nil {
		return err
	}

	return nil
}
