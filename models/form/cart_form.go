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

	if method == "PUT_CART" {
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

	if method == "DELETE" {
		if form.AccessToken == nil {
			return errors.New("Access Token is required")
		}
	}

	if method == "PUT_ITEM" {
		// don't use product_id when update Cart Item
		if form.Product_Id != nil {
			form.Product_Id = nil
		}

		if form.Quantity == nil {
			return errors.New("No Quantity specified")
		}
		if *form.Quantity < 0 {
			return errors.New("Quantity should bigger or equal to 0")
		}
	}

	return nil
}

func (form *CartForm) GenerateToken() string {
	return ""
}

func (form *CartForm) Cart(app *App, item_id ...uint) (*Cart, error) {
	var cart = new(Cart)
	var err error

	if form.AccessToken != nil {
		// Get cart info based on AccessToken
		if cart, err = (CartRepository{app.DB}).GetCart(*form.AccessToken); err != nil {
			if err.Error() == "record not found" {
				return cart, errors.New("Access Token is invalid")
			}

			// unknown err, return anyway
			return cart, err
		}
	}

	// DELETE method should not update
	if form.Product_Id != nil && form.Quantity != nil {
		err = cart.UpdateItems(form.Product_Id, nil, *form.Quantity)
	}

	// PUT CartItem, should retrieve ProductId based on ItemId
	if form.Product_Id == nil && form.Quantity != nil {
		err = cart.UpdateItems(nil, &item_id[0], *form.Quantity)
	}

	return cart, err
}
