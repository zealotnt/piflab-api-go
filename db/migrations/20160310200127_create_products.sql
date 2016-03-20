
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE products (
	id SERIAL PRIMARY KEY,
	name varchar(80),
	price int, 
	provider varchar(80),
	rating float,
	status varchar(80),
	image varchar(80),
	detail varchar(80),	
	created_at date,
	updated_at date
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE products;
