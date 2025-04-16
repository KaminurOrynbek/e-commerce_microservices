package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrInvalidPrice      = errors.New("price must be greater than or equal to 0")
	ErrInvalidStock      = errors.New("stock must be greater than or equal to 0")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrProductNotFound   = errors.New("product not found")
)

type Product struct {
	id          uint64
	name        string
	description string
	price       float64
	stock       int
	categoryID  uint64
	createdAt   time.Time
	updatedAt   time.Time
	isDeleted   bool
}

func NewProduct(name, description string, price float64, stock int, categoryID uint64) (*Product, error) {
	if price < 0 {
		return nil, ErrInvalidPrice
	}
	if stock < 0 {
		return nil, ErrInvalidStock
	}

	now := time.Now()
	return &Product{
		name:        name,
		description: description,
		price:       price,
		stock:       stock,
		categoryID:  categoryID,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

func (p *Product) ID() uint64 {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Description() string {
	return p.description
}

func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) Stock() int {
	return p.stock
}

func (p *Product) CategoryID() uint64 {
	return p.categoryID
}

func (p *Product) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Product) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Product) IsDeleted() bool {
	return p.isDeleted
}

func (p *Product) Update(name, description string, price float64, stock int, categoryID uint64) error {
	if price < 0 {
		return ErrInvalidPrice
	}
	if stock < 0 {
		return ErrInvalidStock
	}

	p.name = name
	p.description = description
	p.price = price
	p.stock = stock
	p.categoryID = categoryID
	p.updatedAt = time.Now()
	return nil
}

func (p *Product) UpdateStock(quantity int) error {
	newStock := p.stock + quantity
	if newStock < 0 {
		return ErrInsufficientStock
	}
	p.stock = newStock
	p.updatedAt = time.Now()
	return nil
}

func (p *Product) Delete() {
	p.isDeleted = true
	p.updatedAt = time.Now()
}

type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id uint64) (*Product, error)
	List(ctx context.Context, categoryID uint64, offset, limit int) ([]*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id uint64) error
}
