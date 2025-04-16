package dto

import (
	"github.com/KaminurOrynbek/e-commerce_microservices/inventory_service/internal/domain"
	"time"
)

type ProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gte=0"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
	CategoryID  uint64  `json:"category_id" binding:"required"`
}

type ProductResponse struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  uint64    `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (r *ProductRequest) ToProduct() (*domain.Product, error) {
	return domain.NewProduct(r.Name, r.Description, r.Price, r.Stock, r.CategoryID)
}

func FromProduct(p *domain.Product) *ProductResponse {
	return &ProductResponse{
		ID:          p.ID(),
		Name:        p.Name(),
		Description: p.Description(),
		Price:       p.Price(),
		Stock:       p.Stock(),
		CategoryID:  p.CategoryID(),
		CreatedAt:   p.CreatedAt(),
		UpdatedAt:   p.UpdatedAt(),
	}
}

type ProductListResponse struct {
	Data []ProductResponse `json:"data"`
	Meta struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	} `json:"meta"`
}
