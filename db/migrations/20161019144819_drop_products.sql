
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE products;
DROP TYPE product_status;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
CREATE TYPE product_status AS ENUM ('available', 'out_of_stock');
CREATE TABLE products (
	id SERIAL PRIMARY KEY,
	name varchar(255),
	price int,
	provider varchar(255),
	rating float,
	status product_status DEFAULT('available'),
	image varchar(255),
	detail TEXT,
	created_at timestamp,
	updated_at timestamp,
	image_updated_at timestamp
);
