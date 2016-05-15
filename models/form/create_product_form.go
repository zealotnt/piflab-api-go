package models

import (
	"errors"
	"github.com/mholt/binding"
	"net/http"

	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
)

type CreateProductForm struct {
	ProductForm
}

func (form *CreateProductForm) FieldMap(req *http.Request) binding.FieldMap {
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
		&form.Image: binding.Field{
			Form: "image",
		},
	}
}

func (form *CreateProductForm) Validate() error {
	if form.Name == nil {
		return errors.New("Name is required")
	}

	if form.Price == nil {
		return errors.New("Price is required")
	}

	if form.Provider == nil {
		return errors.New("Provider is required")
	}

	if form.Rating == nil {
		return errors.New("Rating is required")
	}
	if *form.Rating > float32(5.0) {
		return errors.New("Rating must be less than or equal to 5")
	}

	if form.Status == nil {
		return errors.New("Status is required")
	}
	if !stringInSlice(*form.Status, STATUS_OPTIONS) {
		return errors.New("Status is invalid")
	}

	if form.Image == nil {
		return errors.New("Image is required")
	}
	if !form.isValidImage() {
		return errors.New("Image extension is invalid")
	}

	return nil
}

func (form *CreateProductForm) Product() *Product {
	return &Product{
		Name:      *form.Name,
		Price:     *form.Price,
		Provider:  *form.Provider,
		Rating:    *form.Rating,
		Status:    *form.Status,
		Image:     form.Image.Filename,
		ImageData: form.ImageData(),
	}
}
