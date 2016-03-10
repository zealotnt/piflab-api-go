
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE products (
	Name varchar(80),
	Price int, 
	Provider varchar(80),
	Rating int,
	Status varchar(80),
	ImageUrl varchar(80),
	DetailUrl varchar(80)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE products;
