
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE products ALTER status DROP DEFAULT;
ALTER TABLE products ALTER COLUMN status TYPE product_status USING status::product_status;
ALTER TABLE products ALTER COLUMN status SET DEFAULT('available');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE products ALTER status DROP DEFAULT;
ALTER TABLE products ALTER COLUMN status TYPE varchar;
ALTER TABLE products ALTER COLUMN status SET DEFAULT('available');
