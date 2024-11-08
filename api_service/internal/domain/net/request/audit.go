package request

import "time"

type GetAllLogsFilter struct {
	UserId        *string    `form:"user_id"`
	ProductId     *string    `form:"product_id"`
	OperationType *string    `form:"operation_type"`
	Timestamp     *time.Time `form:"timestamp"`
	Page          int        `form:"page"`
	Limit         int        `form:"limit"`
}
