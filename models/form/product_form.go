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
	Detail   *string               `json:"detail"`
	Image    *multipart.FileHeader `json:"image"`
	Avatar   *multipart.FileHeader `json:"avatar"`
}

var STATUS_OPTIONS = []string{
	"sale",
	"out of stock",
	"available",
}

var VALIDATE_ERROR_MESSAGE = map[string]string{
	"Required_Name":        "Name is required",
	"Required_Price":       "Price is required",
	"Required_Provider":    "Provider is required",
	"Required_Rating":      "Rating is required",
	"Required_Status":      "Status is required",
	"Required_Detail":      "Detail is required",
	"Invalid_Rating_Big":   "Rating must be less than or equal to 5",
	"Invalid_Rating_Small": "Rating must be bigger than or equal to 0",
	"Invalid_Status":       "Status is invalid",
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
		&form.Detail: binding.Field{
			Form: "detail",
		},
		&form.Image: binding.Field{
			Form: "image",
		},
		&form.Avatar: binding.Field{
			Form: "avatar",
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

func (form *ProductForm) AvatarData() []byte {
	if form.Avatar == nil {
		return nil
	}

	fh, err := form.Avatar.Open()
	if err != nil {
		return nil
	}
	defer fh.Close()

	dataBytes := bytes.Buffer{}

	dataBytes.ReadFrom(fh)

	return dataBytes.Bytes()
}

func (form *ProductForm) Product() *Product {
	var product = new(Product)

	if form.Image != nil {
		product.Image = form.Image.Filename
		product.ImageData = form.ImageData()
		product.ImageThumbnailData = (ImageService{}).GetThumbnail(form.Image, 320)
		product.ImageDetailData = (ImageService{}).GetDetail(form.Image, 550)
	}

	if form.Avatar != nil {
		product.Avatar = form.Avatar.Filename
		product.AvatarData = form.AvatarData()
		product.AvatarThumbnailData = (ImageService{}).GetThumbnail(form.Avatar, 320)
		product.AvatarDetailData = (ImageService{}).GetDetail(form.Avatar, 550)
	}

	product.Name = *form.Name
	product.Price = *form.Price
	product.Provider = *form.Provider
	product.Rating = *form.Rating
	product.Status = *form.Status
	product.Detail = *form.Detail

	return product
}
