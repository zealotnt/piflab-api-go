
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE order_items DROP CONSTRAINT "order_items_product_id_fkey";

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM order_items;
ALTER TABLE order_items ADD FOREIGN KEY(product_id) REFERENCES products(id);
