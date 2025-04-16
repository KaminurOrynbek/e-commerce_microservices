package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type Category struct {
	id          uint64
	name        string
	description string
	createdAt   time.Time
	updatedAt   time.Time
}

func NewCategory(name, description string) *Category {
	now := time.Now()
	return &Category{
		name:        name,
		description: description,
		createdAt:   now,
		updatedAt:   now,
	}
}

func (c *Category) ID() uint64 {
	return c.id
}

func (c *Category) SetID(id uint64) {
	c.id = id
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) Description() string {
	return c.description
}

func (c *Category) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Category) UpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Category) Update(name, description string) {
	c.name = name
	c.description = description
	c.updatedAt = time.Now()
}

type CategoryRepository interface {
	Create(ctx context.Context, category *Category) error
	GetByID(ctx context.Context, id uint64) (*Category, error)
	List(ctx context.Context, offset, limit int) ([]*Category, error)
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id uint64) error
}
