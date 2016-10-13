
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE order_status_logs (
	id SERIAL PRIMARY KEY,
	code varchar(255),
	status order_status DEFAULT 'processing',
	created_at timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE order_status_logs;
