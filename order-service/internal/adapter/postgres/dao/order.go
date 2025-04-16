package dao

type OrderedProduct struct {
	ProductID int64 `db:"product_id"`
	Quantity  int   `db:"quantity"`
}

type Order struct {
	ID           int64            `db:"id"`
	UserID       int64            `db:"user_id"`
	Products     []OrderedProduct `db:"products"`
	TotalAmount  float64          `db:"total_amount"`
	Status       string           `db:"status"`
	DeliveryAddr string           `db:"delivery_addr"`
}
