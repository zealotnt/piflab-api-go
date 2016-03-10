
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO products VALUES (
	'XBox',
	70000,
	'Microsoft',
	3,
	'sale',
	'/img/catalog/1.png',
	'/products/1'
);

INSERT INTO products VALUES (
	'PS3',
	50000,
	'Sony',
	4,
	'sale',
	'/img/catalog/2.png',
	'/products/2'
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM products WHERE DetailUrl = '/products/1';
DELETE FROM products WHERE DetailUrl = '/products/2';
