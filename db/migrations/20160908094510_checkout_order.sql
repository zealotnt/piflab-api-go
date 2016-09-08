
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE orders
ADD order_code TEXT NOT NULL UNIQUE,
ADD customer_name varchar(100),
ADD customer_address varchar(255),
ADD customer_phone varchar(20),
ADD customer_email varchar(20),
ADD customer_note varchar(255),
ADD created_at timestamp,
ADD updated_at timestamp;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE orders
DROP COLUMN order_code,
DROP COLUMN customer_name,
DROP COLUMN customer_address,
DROP COLUMN customer_phone,
DROP COLUMN customer_email,
DROP COLUMN customer_note,
DROP COLUMN created_at,
DROP COLUMN updated_at;
