
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE products ADD avatar varchar(255), ADD avatar_updated_at timestamp;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE products DROP COLUMN avatar, DROP COLUMN avatar_updated_at;
