package domain

type OrderedProduct struct {
	ProductID int64
	Quantity  int
}

type Order struct {
	ID           int64
	UserID       int64
	Products     []OrderedProduct
	TotalAmount  float64
	Status       string
	DeliveryAddr string
}
