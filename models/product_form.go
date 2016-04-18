package models

import (
	"bytes"
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

func (form ProductForm) isValidImage() (bool, error) {
	fh, err := form.Image.Open()
	if err != nil {
		return false, err
	}
	defer fh.Close()

	buff := make([]byte, 512)
	if _, err := fh.Read(buff); err != nil {
		return false, err
	}

	filetype := http.DetectContentType(buff)
	switch filetype {
	case "image/jpeg":
		fallthrough
	case "image/png":
		fallthrough
	case "image/gif":
		return true, nil
	default:
		return false, nil
	}

	return false, nil
}

func (form ProductForm) Validate() error {
	if form.Image == nil {
		return errors.New("Image is required")
	}

	if !stringInSlice(form.Status, STATUS_OPTIONS) {
		return errors.New("Status is invalid")
	}

	if valid, err := form.isValidImage(); valid == false || err != nil {
		if err != nil {
			return err
		}
		return errors.New("Invalid image extension")
	}

	return nil
}

func (form *ProductForm) Product() Product {
	return Product{
		Name:      form.Name,
		Price:     form.Price,
		Provider:  form.Provider,
		Rating:    form.Rating,
		Status:    form.Status,
		Image:     form.Image.Filename,
		ImageData: form.ImageBytes(),
	}
}

func (form *ProductForm) ImageBytes() []byte {
	if form.Image == nil {
		return nil
	}

	fh, err := form.Image.Open()
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	dataBytes := bytes.Buffer{}

	dataBytes.ReadFrom(fh)

	return dataBytes.Bytes()
}
