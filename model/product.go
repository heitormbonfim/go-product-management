package model

type Product struct {
	ID    int     `json:"product_id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
