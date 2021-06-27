package models

type Product struct {
	ID          int64  `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

type Purchase struct {
	ID         int64   `json:"id"`
	Product    Product `json:"product_id"`
	Addons     []Addon `json:"addons"`
	TotalPrice int64   `json:"total_price"`
}
