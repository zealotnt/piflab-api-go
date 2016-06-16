package repository

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/o0khoiclub0o/piflab-store-api-go/services"
	"time"
)

type ProductRepository struct {
	*DB
}

func (repo ProductRepository) FindById(id uint) (*Product, error) {
	product := &Product{}

	err := repo.DB.First(&product, id).Error
	if err != nil {
		return nil, err
	}

	err = product.GetImageUrl()

	return product, err
}

func (repo ProductRepository) GetAll() (*[]Product, error) {
	products := &[]Product{}
	err := repo.DB.Find(products).Error

	for idx := range *products {
		(*products)[idx].GetImageUrl()
	}

	return products, err
}

func (repo ProductRepository) saveFile(product *Product) error {
	if err := (FileService{}).SaveFile(product.ImageData, product.GetImagePath(ORIGIN)); err != nil {
		return err
	}

	if err := (FileService{}).SaveFile(product.ImageThumbnailData, product.GetImagePath(THUMBNAIL)); err != nil {
		return err
	}

	if err := (FileService{}).SaveFile(product.ImageDetailData, product.GetImagePath(DETAIL)); err != nil {
		return err
	}
	return nil
}

func (repo ProductRepository) deleteFile(product *Product) error {
	if err := (FileService{}).DeleteFile(product.GetImagePath(ORIGIN)); err != nil {
		return err
	}

	if err := (FileService{}).DeleteFile(product.GetImagePath(THUMBNAIL)); err != nil {
		return err
	}

	if err := (FileService{}).DeleteFile(product.GetImagePath(DETAIL)); err != nil {
		return err
	}
	return nil
}

func (repo ProductRepository) createProduct(product *Product) error {
	product.ImageUpdatedAt = time.Now()

	tx := repo.DB.Begin()

	if err := tx.Create(product).Error; err != nil {
		tx.Rollback()
		return err
	}

	if product.Image != "" {
		if err := repo.saveFile(product); err != nil {
			tx.Rollback()
			return err
		}
		product.GetImageUrl()
	}

	tx.Commit()

	return nil
}

func (repo ProductRepository) updateProduct(product *Product) error {
	tx := repo.DB.Begin()

	if product.ImageData != nil {
		if err := repo.deleteFile(product); err != nil {
			tx.Rollback()
			return err
		}

		product.ImageUpdatedAt = time.Now()

		if err := repo.saveFile(product); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Save(product).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (repo ProductRepository) SaveProduct(product *Product) error {
	if product.Id == 0 {
		return repo.createProduct(product)
	}
	return repo.updateProduct(product)
}

func (repo ProductRepository) CountProduct() (int, error) {
	count := 0

	err := repo.DB.Table("products").Count(&count).Error

	return count, err
}

func (repo ProductRepository) DeleteProduct(id uint) (*Product, error) {
	product, err := repo.FindById(id)
	if err != nil {
		return product, err
	}

	tx := repo.DB.Begin()

	if err := repo.deleteFile(product); err != nil {
		tx.Rollback()
		return product, err
	}

	if err := repo.DB.Delete(product).Error; err != nil {
		tx.Rollback()
		return product, err
	}

	tx.Commit()

	return product, nil
}
