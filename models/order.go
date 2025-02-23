package models

type Order struct {
	ID       int     `json:"id"`
	Symbol   string  `json:"symbol"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Type     string  `json:"type"`
	Status   string  `json:"status"`
}
