package repository

import (
    "database/sql"
    "errors"
    "github.com/KaminurOrynbek/e-commerce_microservices/order-service/internal/domain"
)

type PgOrderRepository struct {
    db *sql.DB
}

func NewPgOrderRepository(db *sql.DB) *PgOrderRepository {
    return &PgOrderRepository{db: db}
}

func (r *PgOrderRepository) CreateOrder(o domain.Order) (domain.Order, error) {
    query := `
        INSERT INTO orders (user_id, total_amount, status, delivery_address)
        VALUES ($1, $2, $3, $4) RETURNING id
    `
    err := r.db.QueryRow(query, o.UserID, o.TotalAmount, o.Status, o.DeliveryAddr).Scan(&o.ID)
    if err != nil {
        return domain.Order{}, err
    }

    // Insert products into a separate table
    for _, product := range o.Products {
        productQuery := `
            INSERT INTO order_products (order_id, product_id, quantity)
            VALUES ($1, $2, $3)
        `
        _, err := r.db.Exec(productQuery, o.ID, product.ProductID, product.Quantity)
        if err != nil {
            return domain.Order{}, err
        }
    }

    return o, nil
}

func (r *PgOrderRepository) GetOrder(id int64) (domain.Order, error) {
    query := `
        SELECT id, user_id, total_amount, status, delivery_address
        FROM orders WHERE id = $1
    `
    var o domain.Order
    err := r.db.QueryRow(query, id).Scan(&o.ID, &o.UserID, &o.TotalAmount, &o.Status, &o.DeliveryAddr)
    if err == sql.ErrNoRows {
        return domain.Order{}, errors.New("order not found")
    } else if err != nil {
        return domain.Order{}, err
    }

    // Fetch products for the order
    productQuery := `
        SELECT product_id, quantity
        FROM order_products WHERE order_id = $1
    `
    rows, err := r.db.Query(productQuery, o.ID)
    if err != nil {
        return domain.Order{}, err
    }
    defer rows.Close()

    var products []domain.OrderedProduct
    for rows.Next() {
        var product domain.OrderedProduct
        if err := rows.Scan(&product.ProductID, &product.Quantity); err != nil {
            return domain.Order{}, err
        }
        products = append(products, product)
    }
    o.Products = products

    return o, nil
}

func (r *PgOrderRepository) UpdateOrder(o domain.Order) (domain.Order, error) {
    query := `
        UPDATE orders
        SET user_id = $1, total_amount = $2, status = $3, delivery_address = $4
        WHERE id = $5
    `
    result, err := r.db.Exec(query, o.UserID, o.TotalAmount, o.Status, o.DeliveryAddr, o.ID)
    if err != nil {
        return domain.Order{}, err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil || rowsAffected == 0 {
        return domain.Order{}, errors.New("order not found")
    }

    // Update products: delete existing and re-insert
    deleteQuery := `DELETE FROM order_products WHERE order_id = $1`
    _, err = r.db.Exec(deleteQuery, o.ID)
    if err != nil {
        return domain.Order{}, err
    }

    for _, product := range o.Products {
        productQuery := `
            INSERT INTO order_products (order_id, product_id, quantity)
            VALUES ($1, $2, $3)
        `
        _, err := r.db.Exec(productQuery, o.ID, product.ProductID, product.Quantity)
        if err != nil {
            return domain.Order{}, err
        }
    }

    return o, nil
}

func (r *PgOrderRepository) ListOrdersByUser(userID int64) ([]domain.Order, error) {
    query := `
        SELECT id, user_id, total_amount, status, delivery_address
        FROM orders WHERE user_id = $1
    `
    rows, err := r.db.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var orders []domain.Order
    for rows.Next() {
        var o domain.Order
        err := rows.Scan(&o.ID, &o.UserID, &o.TotalAmount, &o.Status, &o.DeliveryAddr)
        if err != nil {
            return nil, err
        }

        // Fetch products for each order
        productQuery := `
            SELECT product_id, quantity
            FROM order_products WHERE order_id = $1
        `
        productRows, err := r.db.Query(productQuery, o.ID)
        if err != nil {
            return nil, err
        }
        defer productRows.Close()

        var products []domain.OrderedProduct
        for productRows.Next() {
            var product domain.OrderedProduct
            if err := productRows.Scan(&product.ProductID, &product.Quantity); err != nil {
                return nil, err
            }
            products = append(products, product)
        }
        o.Products = products

        orders = append(orders, o)
    }

    return orders, nil
}