
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE products (
	id SERIAL PRIMARY KEY,
	name varchar(255),
	price int,
	provider varchar(255),
	rating float,
	status varchar(80),
	image varchar(255),
	detail TEXT,
	created_at timestamp,
	updated_at timestamp,
	image_updated_at timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE products;
