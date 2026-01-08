CREATE TABLE items (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    public_id VARCHAR NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    category_id BIGINT NOT NULL,
    supplier_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT fk_items_category
        FOREIGN KEY (category_id) REFERENCES categories(id),
    CONSTRAINT fk_items_supplier
        FOREIGN KEY (supplier_id) REFERENCES suppliers(id)
);