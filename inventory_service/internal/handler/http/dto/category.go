package dto

import (
	"github.com/KaminurOrynbek/e-commerce_microservices/inventory_service/internal/domain"
	"time"
)

type CategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type CategoryResponse struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (r *CategoryRequest) ToCategory() *domain.Category {
	return domain.NewCategory(r.Name, r.Description)
}

func FromCategory(c *domain.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:          c.ID(),
		Name:        c.Name(),
		Description: c.Description(),
		CreatedAt:   c.CreatedAt(),
		UpdatedAt:   c.UpdatedAt(),
	}
}

type CategoryListResponse struct {
	Data []CategoryResponse `json:"data"`
	Meta struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	} `json:"meta"`
}
