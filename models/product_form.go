package models

import (
	"errors"
	"github.com/mholt/binding"
	"mime/multipart"
	"net/http"
)

type ProductForm struct {
	Name     string                `json:"name"`
	Price    int                   `json:"price"`
	Provider string                `json:"provider"`
	Rating   float32               `json:"rating"`
	Status   string                `json:"status"`
	Image    *multipart.FileHeader `json:"image"`
}

func (form *ProductForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&form.Name: binding.Field{
			Form:     "name",
			Required: true,
		},
		&form.Price: binding.Field{
			Form:     "price",
			Required: true,
		},
		&form.Provider: binding.Field{
			Form:     "provider",
			Required: true,
		},
		&form.Rating: binding.Field{
			Form:     "rating",
			Required: true,
		},
		&form.Status: binding.Field{
			Form:     "status",
			Required: true,
		},
		&form.Image: binding.Field{
			Form:     "image",
			Required: true,
		},
	}
}

func (form ProductForm) Validate() error {
	if form.Image == nil {
		return errors.New("Image is required")
	}
	return nil
}

func (form *ProductForm) Product() Product {
	return Product{
		Name:     form.Name,
		Price:    form.Price,
		Provider: form.Provider,
		Rating:   form.Rating,
		Status:   form.Status,
		Image:    form.Image.Filename,
	}
}
