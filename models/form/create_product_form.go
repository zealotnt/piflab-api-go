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
		return errors.New(VALIDATE_ERROR_MESSAGE["Required_Name"])
	}

	if form.Price == nil || *form.Price == 0 {
		return errors.New(VALIDATE_ERROR_MESSAGE["Required_Price"])
	}

	if form.Provider == nil || *form.Provider == "" {
		return errors.New(VALIDATE_ERROR_MESSAGE["Required_Provider"])
	}

	if form.Rating == nil {
		return errors.New(VALIDATE_ERROR_MESSAGE["Required_Rating"])
	}
	if *form.Rating > float32(5.0) {
		return errors.New(VALIDATE_ERROR_MESSAGE["Invalid_Rating_Big"])
	}
	if *form.Rating < float32(0.0) {
		return errors.New(VALIDATE_ERROR_MESSAGE["Invalid_Rating_Small"])
	}

	if form.Status == nil || *form.Status == "" {
		return errors.New(VALIDATE_ERROR_MESSAGE["Required_Status"])
	}
	if !stringInSlice(*form.Status, STATUS_OPTIONS) {
		return errors.New(VALIDATE_ERROR_MESSAGE["Invalid_Status"])
	}

	if form.Detail == nil || *form.Detail == "" {
		return errors.New(VALIDATE_ERROR_MESSAGE["Required_Detail"])
	}

	if form.Image != nil {
		if valid, err := (ImageService{}).IsValidImage(form.Image); valid != true {
			return err
		}
	}

	return nil
}
