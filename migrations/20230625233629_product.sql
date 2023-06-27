-- +goose Up
-- +goose StatementBegin

CREATE TYPE sizes AS (
    width INT,
    height INT,
    length INT
);

CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    size sizes NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE product;
DROP TYPE sizes;

-- +goose StatementEnd
