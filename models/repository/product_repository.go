package repository

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"

	"encoding/json"
	"errors"
	"strconv"
)

type ProductRepository struct {
	*App
}

func (repo ProductRepository) FindById(id uint) (*Product, error) {
	product := &Product{}
	response, body := repo.App.HttpRequest("GET", repo.PRODUCT_SERVICE+"/products/"+strconv.Itoa(int(id)), "")
	if response.Status != "200 OK" {
		return nil, errors.New("Product not found")
	}

	if err := json.Unmarshal([]byte(body), &product); err != nil {
		return nil, err
	}

	return product, nil
}

func (repo ProductRepository) GetPage(offset uint, limit uint, search string) (*ProductPage, error) {
	product_by_page := &ProductPage{}

	response, body := repo.App.HttpRequest("GET",
		repo.PRODUCT_SERVICE+"/products?offset="+
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

func (repo ProductRepository) createProduct(product *Product) error {
	response, body := repo.App.HttpRequest("POST",
		repo.PRODUCT_SERVICE+"/products",
		product)
	if response.Status != "201 Created" {
		return errors.New(body)
	}
	if err := json.Unmarshal([]byte(body), product); err != nil {
		return err
	}

	return nil
}

func (repo ProductRepository) updateProduct(product *Product) error {
	response, body := repo.App.HttpRequest("PUT",
		repo.PRODUCT_SERVICE+"/products/"+strconv.Itoa(int(product.Id)),
		product)
	if response.Status != "200 OK" {
		return errors.New(body)
	}
	if err := json.Unmarshal([]byte(body), product); err != nil {
		return err
	}

	return nil
}

func (repo ProductRepository) SaveProduct(product *Product) error {
	if product.Id == 0 {
		return repo.createProduct(product)
	}
	return repo.updateProduct(product)

	return nil
}

func (repo ProductRepository) DeleteProduct(id uint) (*Product, error) {
	product, err := repo.FindById(id)
	if err != nil {
		return product, err
	}
	response, body := repo.App.HttpRequest("DELETE", repo.PRODUCT_SERVICE+"/products/"+strconv.Itoa(int(id)), "")
	if response.Status != "200 OK" {
		return nil, errors.New(body)
	}

	return product, nil
}
