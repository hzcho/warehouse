package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OperationLog struct {
	OprationId    primitive.ObjectID     `bson:"_id" json:"operation_id"`
	UserId        string                 `bson:"user_id" json:"user_id"`
	ProductId     string                 `bson:"product_id" json:"product_id"`
	OperationType string                 `bson:"operation_type" json:"operation_type"`
	Timestamp     time.Time              `bson:"timestamp" json:"timestamp"`
	UpdatedFields map[string]interface{} `bson:"updated_fields,omitempty" json:"updated_fields,omitempty"`
}
