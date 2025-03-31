package models

type Stock struct {
	StockId int64  `json:"stock_id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Company string `json:"company"`
}
