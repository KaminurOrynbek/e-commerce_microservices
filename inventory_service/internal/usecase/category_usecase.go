package usecase

import (
	"context"
	"github.com/KaminurOrynbek/e-commerce_microservices/inventory_service/internal/domain"
)

type CategoryUseCase struct {
	categoryRepo domain.CategoryRepository
}

func NewCategoryUseCase(repo domain.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		categoryRepo: repo,
	}
}

func (u *CategoryUseCase) CreateCategory(ctx context.Context, name, description string) (*domain.Category, error) {
	category := domain.NewCategory(name, description)
	if err := u.categoryRepo.Create(ctx, category); err != nil {
		return nil, err
	}
	return category, nil
}

func (u *CategoryUseCase) GetCategory(ctx context.Context, id uint64) (*domain.Category, error) {
	category, err := u.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, domain.ErrCategoryNotFound
	}
	return category, nil
}

func (u *CategoryUseCase) ListCategories(ctx context.Context, offset, limit int) ([]*domain.Category, error) {
	if limit <= 0 {
		limit = 10 // Default limit
	}
	if limit > 100 {
		limit = 100 // Max limit
	}
	if offset < 0 {
		offset = 0
	}

	return u.categoryRepo.List(ctx, offset, limit)
}

func (u *CategoryUseCase) UpdateCategory(ctx context.Context, id uint64, name, description string) error {
	category, err := u.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if category == nil {
		return domain.ErrCategoryNotFound
	}

	category.Update(name, description)
	return u.categoryRepo.Update(ctx, category)
}

func (u *CategoryUseCase) DeleteCategory(ctx context.Context, id uint64) error {
	category, err := u.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if category == nil {
		return domain.ErrCategoryNotFound
	}

	return u.categoryRepo.Delete(ctx, id)
}
