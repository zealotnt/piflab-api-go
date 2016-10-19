package repository

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"

	"encoding/json"
	"errors"
	"strconv"
	// "strings"
	// "time"
)

type ProductRepository struct {
	*App
}

func (repo ProductRepository) FindById(id uint) (*Product, error) {
	product := &Product{}
	response, body := repo.App.HttpRequest("GET", "http://product_service:9901/products/"+strconv.Itoa(int(id)), "")
	if response.Status != "200 OK" {
		return nil, errors.New(body)
	}

	if err := json.Unmarshal([]byte(body), &product); err != nil {
		return nil, err
	}

	return product, nil
}

func (repo ProductRepository) GetPage(offset uint, limit uint, search string) (*ProductPage, error) {
	product_by_page := &ProductPage{}

	response, body := repo.App.HttpRequest("GET",
		"http://product_service:9901/products?offset="+
			strconv.Itoa(int(offset))+
			"&limit="+strconv.Itoa(int(limit))+
			"&q="+search,
		"")
	if response.Status != "200 OK" {
		return nil, errors.New(body)
	}

	if err := json.Unmarshal([]byte(body), &product_by_page); err != nil {
		return nil, err
	}

	return product_by_page, nil
}

func (repo ProductRepository) saveFile(product *Product) error {
	// type image_to_save struct {
	// 	data  *[]byte
	// 	field ImageField
	// 	size  ImageSize
	// }

	// if product.Image != "" {
	// 	var images = []image_to_save{
	// 		{&product.ImageData, IMAGE, ORIGIN},
	// 		{&product.ImageThumbnailData, IMAGE, THUMBNAIL},
	// 		{&product.ImageDetailData, IMAGE, DETAIL}}

	// 	for _, image := range images {
	// 		if err := (FileService{}).SaveFile(
	// 			*image.data,
	// 			product.GetImagePath(image.field, image.size),
	// 			product.GetImageContentType(image.field, image.size)); err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	// return nil

	return nil
}

func (repo ProductRepository) deleteFile(product *Product) error {
	// var fields = []ImageField{IMAGE}
	// var sizes = []ImageSize{ORIGIN, THUMBNAIL, DETAIL}

	// for _, field := range fields {
	// 	for _, size := range sizes {
	// 		if err := (FileService{}).DeleteFile(product.GetImagePath(field, size)); err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	// return nil

	return nil
}

func (repo ProductRepository) createProduct(product *Product) error {
	// product.ImageUpdatedAt = time.Now()

	// tx := repo.DB.Begin()

	// if err := tx.Create(product).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// if err := repo.saveFile(product); err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	// product.GetImageUrl()

	// tx.Commit()

	// return nil

	return nil
}

func (repo ProductRepository) updateProduct(product *Product) error {
	// tx := repo.DB.Begin()

	// if product.ImageData != nil {
	// 	if err := repo.deleteFile(product); err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}

	// 	product.Image = product.NewImage
	// 	product.ImageUpdatedAt = time.Now()

	// 	if err := repo.saveFile(product); err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}
	// 	product.GetImageUrl()
	// }

	// if err := tx.Save(product).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// tx.Commit()

	// return nil

	return nil
}

func (repo ProductRepository) SaveProduct(product *Product) error {
	// if product.Id == 0 {
	// 	return repo.createProduct(product)
	// }
	// return repo.updateProduct(product)

	return nil
}

func (repo ProductRepository) CountProduct() (uint, error) {
	// count := uint(0)

	// err := repo.DB.Table("products").Count(&count).Error

	// return count, err

	return 0, nil
}

func (repo ProductRepository) DeleteProduct(id uint) (*Product, error) {
	// // product, err := repo.FindById(id)1
	// product := &Product{}
	// var err error
	// if err != nil {
	// 	return product, err
	// }

	// tx := repo.DB.Begin()

	// if err := repo.deleteFile(product); err != nil {
	// 	tx.Rollback()
	// 	return product, err
	// }

	// if err := repo.DB.Delete(product).Error; err != nil {
	// 	tx.Rollback()
	// 	return product, err
	// }

	// tx.Commit()

	// return product, nil

	return nil, nil
}
