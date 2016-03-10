package repository

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
)

type ProductRepository struct {
	*DB
}

func (repo ProductRepository) GetAll() (*[]Product, error) {
	products := &[]Product{}
	repo.DB.Find(products)
	return products, nil
}