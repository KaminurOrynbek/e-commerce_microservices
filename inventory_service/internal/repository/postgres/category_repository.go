package postgres

import (
	"context"
	"database/sql"
	"github.com/KaminurOrynbek/e-commerce_microservices/inventory_service/internal/domain"
)

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) domain.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(ctx context.Context, category *domain.Category) error {
	query := `
		INSERT INTO categories (name, description, created_at, updated_at)
		VALUES ($1, $2, NOW(), NOW())
		RETURNING id, created_at, updated_at`

	var id uint64
	var createdAt, updatedAt string
	err := r.db.QueryRowContext(
		ctx,
		query,
		category.Name(),
		category.Description(),
	).Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		return err
	}

	category.SetID(id)
	return nil
}

func (r *categoryRepository) GetByID(ctx context.Context, id uint64) (*domain.Category, error) {
	category := &domain.Category{}

	query := `
		SELECT id, name, description, created_at, updated_at
		FROM categories
		WHERE id = $1 AND is_deleted = false`

	var createdAt, updatedAt string
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		func(id uint64) error {
			category.SetID(id)
			return nil
		},
		func(name string) error {
			category.Update(name, category.Description())
			return nil
		},
		func(description string) error {
			category.Update(category.Name(), description)
			return nil
		},
		&createdAt,
		&updatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoryRepository) List(ctx context.Context, offset, limit int) ([]*domain.Category, error) {
	query := `
		SELECT id, name, description, created_at, updated_at
		FROM categories
		WHERE is_deleted = false
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*domain.Category
	for rows.Next() {
		category := &domain.Category{}
		var createdAt, updatedAt string
		err := rows.Scan(
			func(id uint64) error {
				category.SetID(id)
				return nil
			},
			func(name string) error {
				category.Update(name, category.Description())
				return nil
			},
			func(description string) error {
				category.Update(category.Name(), description)
				return nil
			},
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, rows.Err()
}

func (r *categoryRepository) Update(ctx context.Context, category *domain.Category) error {
	query := `
		UPDATE categories
		SET name = $1, description = $2, updated_at = NOW()
		WHERE id = $3 AND is_deleted = false
		RETURNING updated_at`

	var updatedAt string
	err := r.db.QueryRowContext(
		ctx,
		query,
		category.Name(),
		category.Description(),
		category.ID(),
	).Scan(&updatedAt)

	if err == sql.ErrNoRows {
		return sql.ErrNoRows
	}
	if err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) Delete(ctx context.Context, id uint64) error {
	query := `
		UPDATE categories
		SET is_deleted = true, updated_at = NOW()
		WHERE id = $1 AND is_deleted = false`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
