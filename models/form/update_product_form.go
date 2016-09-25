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
		&form.Avatar: binding.Field{
			Form: "avatar",
		},
		&form.Fields: binding.Field{
			Form: "fields",
		},
	}
}

func (form *UpdateProductForm) Validate() error {
	if form.Name != nil {
		if *form.Name == "" {
			return errors.New(VALIDATE_ERROR_MESSAGE["Required_Name"])
		}
	}

	if form.Provider != nil {
		if *form.Provider == "" {
			return errors.New(VALIDATE_ERROR_MESSAGE["Required_Provider"])
		}
	}

	if form.Rating != nil {
		if *form.Rating > float32(5.0) {
			return errors.New(VALIDATE_ERROR_MESSAGE["Invalid_Rating_Big"])
		}
		if *form.Rating < float32(0.0) {
			return errors.New(VALIDATE_ERROR_MESSAGE["Invalid_Rating_Small"])
		}
	}

	if form.Status != nil {
		if *form.Status == "" {
			return errors.New(VALIDATE_ERROR_MESSAGE["Required_Status"])
		}
		if !stringInSlice(*form.Status, STATUS_OPTIONS) {
			return errors.New(VALIDATE_ERROR_MESSAGE["Invalid_Status"])
		}
	}

	if form.Detail != nil {
		if *form.Detail == "" {
			return errors.New(VALIDATE_ERROR_MESSAGE["Required_Detail"])
		}
	}

	if form.Image != nil {
		if valid, err := (ImageService{}).IsValidImage(form.Image); valid != true {
			return err
		}
	}

	if form.Avatar != nil {
		if valid, err := (ImageService{}).IsValidImage(form.Avatar); valid != true {
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
		product.NewImage = form.Image.Filename
		product.ImageData = form.ImageData()
		product.ImageThumbnailData = (ImageService{}).GetThumbnail(form.Image, 320)
		product.ImageDetailData = (ImageService{}).GetDetail(form.Image, 550)
	}

	if form.Avatar != nil {
		product.NewAvatar = form.Avatar.Filename
		product.AvatarData = form.AvatarData()
		product.AvatarThumbnailData = (ImageService{}).GetThumbnail(form.Avatar, 320)
		product.AvatarDetailData = (ImageService{}).GetDetail(form.Avatar, 550)
	}
}
