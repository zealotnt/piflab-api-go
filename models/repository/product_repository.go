package repository

import (
	"errors"
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/o0khoiclub0o/piflab-store-api-go/services"
	"time"
)

type ProductRepository struct {
	*DB
}

func (repo ProductRepository) GetAll() (*[]Product, error) {
	products := &[]Product{}
	err := repo.DB.Find(products).Error
	return products, err
}

func (repo ProductRepository) SaveProduct(product *Product) error {
	if product.ImageData == nil {
		return errors.New("ImageData is required")
	}

	product.ImageUpdatedAt = time.Now()

	tx := repo.DB.Begin()

	if err := tx.Create(product).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := (FileService{}).SaveFile(product.ImageData, product.GetImagePath()); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (repo ProductRepository) CountProduct() (int, error) {
	count := 0

	err := repo.DB.Table("products").Count(&count).Error

	return count, err
}
