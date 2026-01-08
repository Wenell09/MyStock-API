CREATE TABLE stock_transactions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    public_id VARCHAR NOT NULL UNIQUE,
    item_id BIGINT NOT NULL,
    warehouse_id BIGINT NOT NULL,
    type stock_transaction_type NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    created_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT fk_stock_item
        FOREIGN KEY (item_id) REFERENCES items(id),
    CONSTRAINT fk_stock_warehouse
        FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);