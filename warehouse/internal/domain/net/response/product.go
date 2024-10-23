package response

import "warehouse/internal/domain/model"

type Products struct {
	Page     int
	Limit    int
	products []model.Product
}
