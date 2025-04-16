package usecase

import (
	"github.com/KaminurOrynbek/e-commerce_microservices/order-service/internal/domain"
	"github.com/KaminurOrynbek/e-commerce_microservices/order-service/internal/repository"
)

// OrderUseCase defines the methods available for order business logic.
type OrderUseCase interface {
	CreateOrder(o domain.Order) (domain.Order, error)
	GetOrder(id int64) (domain.Order, error)
	UpdateOrder(o domain.Order) (domain.Order, error)
	ListOrdersByUser(userID int64) ([]domain.Order, error)
}

type orderUseCase struct {
	repo *repository.PgOrderRepository
}

// NewOrderUseCase creates a new instance of the order use case.
func NewOrderUseCase(repo *repository.PgOrderRepository) OrderUseCase {
	return &orderUseCase{repo: repo}
}

func (u *orderUseCase) CreateOrder(o domain.Order) (domain.Order, error) {
	return u.repo.CreateOrder(o)
}

func (u *orderUseCase) GetOrder(id int64) (domain.Order, error) {
	return u.repo.GetOrder(id)
}

func (u *orderUseCase) UpdateOrder(o domain.Order) (domain.Order, error) {
	return u.repo.UpdateOrder(o)
}

func (u *orderUseCase) ListOrdersByUser(userID int64) ([]domain.Order, error) {
	return u.repo.ListOrdersByUser(userID)
}
