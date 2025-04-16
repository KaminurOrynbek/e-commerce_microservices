CREATE TABLE IF NOT EXISTS orders
(
    id               SERIAL PRIMARY KEY,
    user_id          INTEGER        NOT NULL,
    products         JSONB          NOT NULL, -- Сериализованный список товаров
    total_amount     NUMERIC(10, 2) NOT NULL,
    status           VARCHAR(50)    NOT NULL,
    delivery_address VARCHAR(255)   NOT NULL,
    created_at       TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
