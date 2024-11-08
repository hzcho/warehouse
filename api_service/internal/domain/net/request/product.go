package request

import (
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateProduct struct {
	Name          string                  `json:"name" form:"name" bson:"name"`
	Description   string                  `json:"description" form:"description" bson:"description"`
	CategoryName  string                  `json:"category_name" form:"category_name" bson:"category_name"`
	Price         float64                 `json:"price" form:"price" bson:"price"`
	StockLevel    int                     `json:"stock_level" form:"stock_level" bson:"stock_level"`
	MinStockLevel int                     `json:"min_stock_level" form:"min_stock_level" bson:"min_stock_level"`
	Manufacturer  string                  `json:"manufacturer" form:"manufacturer" bson:"manufacturer"`
	Supplier      string                  `json:"supplier" form:"supplier" bson:"supplier"`
	Weight        float64                 `json:"weight" form:"weight" bson:"weight"`
	Dimensions    Dimensions              `json:"dimensions" form:"dimensions" bson:"dimensions"`
	Images        []*multipart.FileHeader `form:"images"`
	Tags          []string                `json:"tags" form:"tags[]" bson:"tags"`
}

type UpdateUser struct {
	ID            *primitive.ObjectID      `bson:"_id"`
	Name          *string                  `json:"name" form:"name" bson:"name"`
	Description   *string                  `json:"description" form:"description" bson:"description"`
	CategoryName  *string                  `json:"category_name" form:"category_name" bson:"category_name"`
	Price         *float64                 `json:"price" form:"price" bson:"price"`
	StockLevel    *int                     `json:"stock_level" form:"stock_level" bson:"stock_level"`
	MinStockLevel *int                     `json:"min_stock_level" form:"min_stock_level" bson:"min_stock_level"`
	Manufacturer  *string                  `json:"manufacturer" form:"manufacturer" bson:"manufacturer"`
	Supplier      *string                  `json:"supplier" form:"supplier" bson:"supplier"`
	Weight        *float64                 `json:"weight" form:"weight" bson:"weight"`
	Dimensions    *Dimensions              `json:"dimensions" form:"dimensions" bson:"dimensions"`
	Tags          *[]string                `json:"tags" form:"tags[]" bson:"tags"`
	ImageURLs     *[]string                `json:"image_urls" form:"image_urls"`
	Images        *[]*multipart.FileHeader `form:"images[]" bson:"images"`
	IsActive      *bool                    `json:"is_active" bson:"is_active"`
}

type Dimensions struct {
	Length float64 `json:"length" form:"length" bson:"length"`
	Width  float64 `json:"width" form:"width" bson:"width"`
	Height float64 `json:"height" form:"height" bson:"height"`
}

type GetAllFilter struct {
	ProductName *string `json:"product_name"`
	Page        *int    `json:"page"`
	Limit       *int    `json:"limit"`
}

type UpdateStockLevel struct {
	Id         string `json:"id"`
	StockLevel int    `json:"stock_level"`
}
