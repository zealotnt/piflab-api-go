
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- +goose StatementBegin
CREATE TABLE carts (
	id SERIAL PRIMARY KEY,
	access_token TEXT NOT NULL,
	status varchar(50)
);
CREATE TABLE cart_items (
	id SERIAL PRIMARY KEY,
	cart_id SERIAL REFERENCES carts(id),
	product_id int NOT NULL REFERENCES products(id),
	quantity int
);
-- +goose StatementEnd

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +goose StatementBegin
DROP TABLE cart_items;
DROP TABLE carts;
-- +goose StatementEnd
