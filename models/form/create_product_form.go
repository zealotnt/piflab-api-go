package models

import (
	"errors"
	"github.com/mholt/binding"
	"net/http"

	. "github.com/o0khoiclub0o/piflab-store-api-go/services"
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
		&form.Detail: binding.Field{
			Form: "detail",
		},
		&form.Image: binding.Field{
			Form: "image",
		},
	}
}

func (form *CreateProductForm) Validate() error {
	if form.Name == nil || *form.Name == "" {
		return errors.New("Name is required")
	}

	if form.Price == nil || *form.Price == 0 {
		return errors.New("Price is required")
	}

	if form.Provider == nil || *form.Provider == "" {
		return errors.New("Provider is required")
	}

	if form.Rating == nil {
		return errors.New("Rating is required")
	}
	if *form.Rating > float32(5.0) {
		return errors.New("Rating must be less than or equal to 5")
	}
	if *form.Rating == float32(0.0) {
		return errors.New("Rating must be bigger than 0")
	}

	if form.Status == nil || *form.Status == "" {
		return errors.New("Status is required")
	}
	if !stringInSlice(*form.Status, STATUS_OPTIONS) {
		return errors.New("Status is invalid")
	}

	if form.Detail == nil || *form.Detail == "" {
		return errors.New("Detail is required")
	}

	if form.Image != nil {
		if valid, err := (ImageService{}).IsValidImage(form.Image); valid != true {
			return err
		}
	}

	return nil
}
