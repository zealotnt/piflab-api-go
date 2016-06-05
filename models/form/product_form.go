package models

import (
	"github.com/mholt/binding"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/o0khoiclub0o/piflab-store-api-go/services"

	"bytes"
	"mime/multipart"
	"net/http"
)

type ProductForm struct {
	Name     *string               `json:"name"`
	Price    *int                  `json:"price"`
	Provider *string               `json:"provider"`
	Rating   *float32              `json:"rating"`
	Status   *string               `json:"status"`
	Image    *multipart.FileHeader `json:"image"`
}

var STATUS_OPTIONS = []string{
	"sale",
	"out of stock",
	"available",
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (form *ProductForm) FieldMap(req *http.Request) binding.FieldMap {
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

func (form *ProductForm) ImageData() []byte {
	if form.Image == nil {
		return nil
	}

	fh, err := form.Image.Open()
	if err != nil {
		return nil
	}
	defer fh.Close()

	dataBytes := bytes.Buffer{}

	dataBytes.ReadFrom(fh)

	return dataBytes.Bytes()
}

func (form *ProductForm) Product() *Product {
	return &Product{
		Name:               *form.Name,
		Price:              *form.Price,
		Provider:           *form.Provider,
		Rating:             *form.Rating,
		Status:             *form.Status,
		Image:              form.Image.Filename,
		ImageData:          form.ImageData(),
		ImageThumbnailData: (ImageService{}).GetThumbnail(form.Image, 320),
		ImageDetailData:    (ImageService{}).GetDetail(form.Image, 550),
	}
}
