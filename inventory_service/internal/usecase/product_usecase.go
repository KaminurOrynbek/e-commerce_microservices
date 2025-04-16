package usecase

import (
	"context"
	"github.com/KaminurOrynbek/e-commerce_microservices/inventory_service/internal/domain"
)

type ProductUseCase struct {
	productRepo domain.ProductRepository
}

func NewProductUseCase(repo domain.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepo: repo,
	}
}

func (u *ProductUseCase) CreateProduct(ctx context.Context, product *domain.Product) error {
	// Add business logic validation here
	if product.Price() < 0 {
		return domain.ErrInvalidPrice
	}
	if product.Stock() < 0 {
		return domain.ErrInvalidStock
	}

	return u.productRepo.Create(ctx, product)
}

func (u *ProductUseCase) GetProduct(ctx context.Context, id uint64) (*domain.Product, error) {
	return u.productRepo.GetByID(ctx, id)
}
func (u *ProductUseCase) ListProducts(ctx context.Context, categoryID uint64, offset, limit int) ([]*domain.Product, error) {
	if limit <= 0 {
		limit = 10 // Default limit
	}
	if limit > 100 {
		limit = 100 // Max limit
	}
	if offset < 0 {
		offset = 0
	}

	return u.productRepo.List(ctx, categoryID, offset, limit)
}

func (u *ProductUseCase) UpdateProduct(ctx context.Context, product *domain.Product) error {
	// Add business logic validation here
	if product.Price() < 0 {
		return domain.ErrInvalidPrice
	}
	if product.Stock() < 0 {
		return domain.ErrInvalidStock
	}

	return u.productRepo.Update(ctx, product)
}

func (u *ProductUseCase) DeleteProduct(ctx context.Context, id uint64) error {
	return u.productRepo.Delete(ctx, id)
}

func (u *ProductUseCase) UpdateStock(ctx context.Context, id uint64, quantity int) error {
	product, err := u.productRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if product == nil {
		return domain.ErrProductNotFound
	}

	if err := product.UpdateStock(quantity); err != nil {
		return err
	}

	return u.productRepo.Update(ctx, product)
}
