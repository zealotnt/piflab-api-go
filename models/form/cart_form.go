package models

import (
	"github.com/mholt/binding"
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"

	"errors"
	"net/http"
)

type CartForm struct {
	Product_Id  *uint   `json:"product_id"`
	Quantity    *int    `json:"quantity"`
	AccessToken *string `json:"access_token"`
}

func (form *CartForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&form.Product_Id: binding.Field{
			Form: "product_id",
		},
		&form.Quantity: binding.Field{
			Form: "quantity",
		},
		&form.AccessToken: binding.Field{
			Form: "access_token",
		},
	}
}

func (form *CartForm) Validate(method string) error {
	if method == "GET" {
		if form.AccessToken == nil {
			return errors.New("Access Token is required")
		}
	}

	if method == "PUT" {
		if form.Product_Id == nil {
			return errors.New("No Product selected")
		}

		if form.Quantity == nil {
			return errors.New("No Quantity specified")
		}
		if *form.Quantity == 0 {
			return errors.New("Quantity should not be 0")
		}
	}

	return nil
}

func (form *CartForm) GenerateToken() string {
	return ""
}

func (form *CartForm) Cart(app *App) (*Cart, error) {
	var cart = new(Cart)
	var err error

	if form.AccessToken != nil {
		if cart, err = (CartRepository{app.DB}).GetCart(*form.AccessToken); err != nil {
			if err.Error() == "record not found" {
				return cart, errors.New("Access Token is invalid")
			}
			return cart, err
		}
	}

	err = cart.UpdateItems(*form.Product_Id, *form.Quantity)

	return cart, err
}
