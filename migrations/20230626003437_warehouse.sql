-- +goose Up
-- +goose StatementBegin

CREATE TABLE warehouse (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    availability BOOLEAN NOT NULL
);

CREATE TABLE warehouse_product (
    id SERIAL PRIMARY KEY,
    quantity INT NOT NULL,
    reserved_number INT NOT NULL,
    warehouse_id INT REFERENCES warehouse(id),
    product_id INT REFERENCES product(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE warehouse_product;
DROP TABLE warehouse;

-- +goose StatementEnd
