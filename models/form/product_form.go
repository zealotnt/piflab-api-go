package models

import (
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

func (form ProductForm) isValidImage() bool {
	fh, err := form.Image.Open()
	if err != nil {
		return false
	}
	defer fh.Close()

	buff := make([]byte, 512)
	if _, err := fh.Read(buff); err != nil {
		return false
	}

	filetype := http.DetectContentType(buff)
	switch filetype {
	case "image/jpeg":
		fallthrough
	case "image/png":
		fallthrough
	case "image/gif":
		return true
	default:
		return false
	}

	return false
}

func (form *ProductForm) ImageData() []byte {
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
