package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID            *primitive.ObjectID `json:"id" bson:"_id"`
	Name          *string             `json:"name" bson:"name"`
	Description   *string             `json:"description" bson:"description"`
	CategoryID    *string             `json:"category_id" bson:"category_id"` // Ссылка на категорию
	Price         *float64            `json:"price" bson:"price"`
	StockLevel    *int                `json:"stock_level" bson:"stock_level"`         // Текущий уровень запасов
	MinStockLevel *int                `json:"min_stock_level" bson:"min_stock_level"` // Минимальный уровень запасов
	Manufacturer  *string             `json:"manufacturer" bson:"manufacturer"`       // Производитель
	Supplier      *string             `json:"supplier" bson:"supplier"`               // Поставщик
	Weight        *float64            `json:"weight" bson:"weight"`                   // Вес продукта (например, в кг)
	Dimensions    *Dimensions         `json:"dimensions" bson:"dimensions"`           // Габариты продукта
	CreatedAt     *time.Time          `json:"created_at" bson:"created_at"`           // Дата добавления продукта
	UpdatedAt     *time.Time          `json:"updated_at" bson:"updated_at"`           // Дата последнего обновления
	ImageURLs     *[]string           `json:"images" bson:"images"`                   // Ссылки на изображения продукта
	Tags          *[]string           `json:"tags" bson:"tags"`                       // Теги для фильтрации или поиска
	IsActive      *bool               `json:"is_active" bson:"is_active"`             // Статус продукта (активный/неактивный)
}

type Dimensions struct {
	Length float64 `json:"length" form:"length" bson:"length"`
	Width  float64 `json:"width" form:"width" bson:"width"`
	Height float64 `json:"height" form:"height" bson:"height"`
}
