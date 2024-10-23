package handler

import (
	"context"
	"encoding/json"
	"notification/internal/domain/model"
	"notification/internal/domain/usecase"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Warehouse struct {
	warehouseUseCase usecase.Warehouse
}

func NewWarehouse(warehouseUseCase usecase.Warehouse) *Warehouse {
	return &Warehouse{
		warehouseUseCase: warehouseUseCase,
	}
}

func (h *Warehouse) MinValue(ctx context.Context, msg *kafka.Message) error {
	event := model.MinValue{}
	if err := json.Unmarshal(msg.Value, &event); err != nil {
		return err
	}

	if err := h.warehouseUseCase.MinValue(ctx, event); err != nil {
		return err
	}

	return nil
}
