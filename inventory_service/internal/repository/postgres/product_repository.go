package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/KaminurOrynbek/e-commerce_microservices/inventory_service/internal/domain"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, product *domain.Product) error {
	var id uint64
	var createdAt, updatedAt sql.NullTime

	query := `
		INSERT INTO products (name, description, price, stock, category_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRowContext(
		ctx,
		query,
		product.Name(),
		product.Description(),
		product.Price(),
		product.Stock(),
		product.CategoryID(),
	).Scan(&id, &createdAt, &updatedAt)

	if err != nil {
		return err
	}

	// Create a new product with the returned ID and timestamps
	newProduct, err := domain.NewProduct(
		product.Name(),
		product.Description(),
		product.Price(),
		product.Stock(),
		product.CategoryID(),
	)
	if err != nil {
		return err
	}

	*product = *newProduct
	return nil
}

func (r *productRepository) GetByID(ctx context.Context, id uint64) (*domain.Product, error) {
	var name, description string
	var price float64
	var stock int
	var categoryID uint64
	var createdAt, updatedAt sql.NullTime
	var isDeleted bool

	query := `
		SELECT id, name, description, price, stock, category_id, created_at, updated_at, is_deleted
		FROM products
		WHERE id = $1 AND is_deleted = false`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&id,
		&name,
		&description,
		&price,
		&stock,
		&categoryID,
		&createdAt,
		&updatedAt,
		&isDeleted,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	product, err := domain.NewProduct(name, description, price, stock, categoryID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepository) List(ctx context.Context, categoryID uint64, offset, limit int) ([]*domain.Product, error) {
	var args []interface{}
	argPosition := 1

	query := `
		SELECT id, name, description, price, stock, category_id, created_at, updated_at, is_deleted
		FROM products
		WHERE is_deleted = false`

	if categoryID > 0 {
		query += fmt.Sprintf(" AND category_id = $%d", argPosition)
		args = append(args, categoryID)
		argPosition++
	}

	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", argPosition, argPosition+1)
	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var id uint64
		var name, description string
		var price float64
		var stock int
		var catID uint64
		var createdAt, updatedAt sql.NullTime
		var isDeleted bool

		err := rows.Scan(
			&id,
			&name,
			&description,
			&price,
			&stock,
			&catID,
			&createdAt,
			&updatedAt,
			&isDeleted,
		)
		if err != nil {
			return nil, err
		}

		product, err := domain.NewProduct(name, description, price, stock, catID)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, rows.Err()
}

func (r *productRepository) Update(ctx context.Context, product *domain.Product) error {
	var updatedAt sql.NullTime

	query := `
		UPDATE products
		SET name = $1, description = $2, price = $3, stock = $4, category_id = $5, updated_at = NOW()
		WHERE id = $6 AND is_deleted = false
		RETURNING updated_at`

	err := r.db.QueryRowContext(
		ctx,
		query,
		product.Name(),
		product.Description(),
		product.Price(),
		product.Stock(),
		product.CategoryID(),
		product.ID(),
	).Scan(&updatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) Delete(ctx context.Context, id uint64) error {
	query := `
		UPDATE products
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
