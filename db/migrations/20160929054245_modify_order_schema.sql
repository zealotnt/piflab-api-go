
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE orders ALTER COLUMN customer_phone TYPE VARCHAR(255);
ALTER TABLE orders ALTER COLUMN customer_email TYPE VARCHAR(255);
ALTER TABLE orders ALTER COLUMN customer_name TYPE VARCHAR(255);
ALTER TABLE orders RENAME customer_note TO note;
ALTER TABLE orders RENAME order_code TO code;

ALTER TABLE order_items
ADD product_name varchar(255) NOT NULL DEFAULT 'not save yet',
ADD product_price int NOT NULL DEFAULT 0;
ALTER TABLE order_items ALTER COLUMN product_name DROP DEFAULT;
ALTER TABLE order_items ALTER COLUMN product_price DROP DEFAULT;

UPDATE products SET status='out_of_stock' WHERE status='out of stock';
UPDATE products SET status='available' WHERE status='sale';
CREATE TYPE product_status AS ENUM ('available', 'out_of_stock');
ALTER TABLE products ALTER COLUMN status TYPE varchar USING status::product_status;
ALTER TABLE products ALTER COLUMN status SET DEFAULT('available');

ALTER TABLE products DROP COLUMN avatar, DROP COLUMN avatar_updated_at;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE orders RENAME note TO customer_note;
ALTER TABLE orders RENAME code TO order_code;

ALTER TABLE order_items
DROP COLUMN product_name,
DROP COLUMN product_price;

ALTER TABLE products ALTER COLUMN status TYPE varchar(255);
DROP TYPE product_status;

ALTER TABLE products ADD avatar varchar(255), ADD avatar_updated_at timestamp;
