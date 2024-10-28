package model

type MinValue struct {
	ProductName   string `json:"product_name"`
	StockLevel    int    `json:"stock_level"`
	MinStockLevel int    `json:"min_stock_level"`
}
