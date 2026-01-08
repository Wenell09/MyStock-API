CREATE TABLE item_warehouses (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    public_id VARCHAR NOT NULL UNIQUE,
    item_id BIGINT NOT NULL,
    warehouse_id BIGINT NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT fk_item_warehouses_item
        FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE,
    CONSTRAINT fk_item_warehouses_warehouse
        FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE,
    CONSTRAINT uq_item_warehouse UNIQUE (item_id, warehouse_id)
);