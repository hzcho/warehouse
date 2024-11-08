package response

import "warehouse/internal/domain/model"

type Products struct {
	Page     int             `json:"page"`
	Limit    int             `json:"limit"`
	Products []model.Product `json:"products"`
}
