package response

import "audit/internal/domain/model"

type GetAllResponse struct {
	Page  int                  `json:"page"`
	Limit int                  `json:"limit"`
	Logs  []model.OperationLog `json:"logs"`
}
