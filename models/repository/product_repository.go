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
	err := repo.DB.Find(products).Error
	return products, err
}

func (repo ProductRepository) CreateProduct(product *Product) error {
	err := repo.DB.Create(product).Error
	return err
}
