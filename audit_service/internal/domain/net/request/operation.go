package request

import "time"

type GetAllFilter struct {
	UserId        *string    `json:"user_id"`
	ProductId     *string    `json:"product_id"`
	OperationType *string    `josn:"operation_type"`
	Timestamp     *time.Time `json:"timestamp"`
	Page          *int       `json:"page"`
	Limit         *int       `json:"limit"`
}
