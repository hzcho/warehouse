package model

import (
	"time"
)

type MinValue struct {
	ProductName   string `json:"product_name"`
	StockLevel    int    `json:"stock_level"`
	MinStockLevel int    `json:"min_stock_level"`
}

type OperationLog struct {
	UserId        string                 `bson:"user_id" json:"user_id"`
	ProductId     string                 `bson:"product_id" json:"product_id"`
	OperationType string                 `bson:"operation_type" json:"operation_type"`
	Timestamp     time.Time              `bson:"timestamp" json:"timestamp"`
	UpdatedFields map[string]interface{} `bson:"updated_fields,omitempty" json:"updated_fields,omitempty"`
}
