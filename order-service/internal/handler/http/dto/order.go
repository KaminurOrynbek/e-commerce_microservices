package dto

type OrderedProduct struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

type Order struct {
	ID           int64            `json:"id"`
	UserID       int64            `json:"user_id"`
	Products     []OrderedProduct `json:"products"`
	TotalAmount  float64          `json:"total_amount"`
	Status       string           `json:"status"`
	DeliveryAddr string           `json:"delivery_address"`
}
