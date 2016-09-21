
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
DELETE FROM order_items;
DELETE FROM orders;
ALTER TABLE orders ALTER COLUMN status TYPE varchar(255);
ALTER TABLE orders ALTER COLUMN status SET DEFAULT('temp default');
DROP TYPE order_status;
CREATE TYPE order_status AS ENUM ('cart', 'processing', 'shipping', 'cancelled', 'completed');
ALTER TABLE orders ALTER status DROP DEFAULT;
ALTER TABLE orders ALTER COLUMN status TYPE order_status USING status::order_status;
ALTER TABLE orders ALTER COLUMN status SET DEFAULT('cart');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM order_items;
DELETE FROM orders;
ALTER TABLE orders ALTER COLUMN status TYPE varchar(255);
ALTER TABLE orders ALTER COLUMN status SET DEFAULT('temp default');
DROP TYPE order_status;
CREATE TYPE order_status AS ENUM ('cart', 'processing', 'canceled', 'done');
ALTER TABLE orders ALTER status DROP DEFAULT;
ALTER TABLE orders ALTER COLUMN status TYPE order_status USING status::order_status;
ALTER TABLE orders ALTER COLUMN status SET DEFAULT('cart');
