package models

import (
	"errors"
	"github.com/mholt/binding"
	"net/http"

	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/o0khoiclub0o/piflab-store-api-go/services"
)

type UpdateProductForm struct {
	ProductForm
}

func (form *UpdateProductForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&form.Name: binding.Field{
			Form: "name",
		},
		&form.Price: binding.Field{
			Form: "price",
		},
		&form.Provider: binding.Field{
			Form: "provider",
		},
		&form.Rating: binding.Field{
			Form: "rating",
		},
		&form.Status: binding.Field{
			Form: "status",
		},
		&form.Detail: binding.Field{
			Form: "detail",
		},
		&form.Image: binding.Field{
			Form: "image",
		},
	}
}

func (form *UpdateProductForm) Validate() error {
	if form.Rating != nil {
		if *form.Rating > float32(5.0) {
			return errors.New("Rating must be less than or equal to 5")
		}
		if *form.Rating == float32(0.0) {
			return errors.New("Rating must be bigger than 0")
		}
	}

	if form.Status != nil {
		if !stringInSlice(*form.Status, STATUS_OPTIONS) {
			return errors.New("Status is invalid")
		}
	}

	if form.Image != nil {
		if valid, err := (ImageService{}).IsValidImage(form.Image); valid != true {
			return err
		}
	}

	return nil
}

func (form *UpdateProductForm) Assign(product *Product) {
	if form.Name != nil {
		product.Name = *form.Name
	}

	if form.Price != nil {
		product.Price = *form.Price
	}

	if form.Provider != nil {
		product.Provider = *form.Provider
	}

	if form.Rating != nil {
		product.Rating = *form.Rating
	}

	if form.Status != nil {
		product.Status = *form.Status
	}

	if form.Detail != nil {
		product.Detail = *form.Detail
	}

	if form.Image != nil {
		product.Image = form.Image.Filename
		product.ImageData = form.ImageData()
	}
}
