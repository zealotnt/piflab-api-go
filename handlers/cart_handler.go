package handlers

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/form"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"

	"net/http"
)

func GetCartHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		form := new(CartForm)

		if err := Bind(form, r); err != nil {
			JSON(w, err, 400)
			return
		}

		if err := form.Validate("GET"); err != nil {
			JSON(w, err, 401)
			return
		}

		cart, err := (CartRepository{app.DB}).GetCart(*form.AccessToken)
		if err != nil {
			JSON(w, err, 500)
			return
		}
		cart.CalculateAmount()

		JSON(w, cart)
	}
}

func UpdateCartHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		form := new(CartForm)

		if err := Bind(form, r); err != nil {
			JSON(w, err, 400)
		}

		if err := form.Validate("PUT_CART"); err != nil {
			JSON(w, err, 422)
			return
		}

		cart, err := form.Cart(app)
		if err != nil {
			JSON(w, err, 422)
			return
		}
		if err := (CartRepository{app.DB}).SaveCart(cart); err != nil {
			JSON(w, err, 500)
			return
		}

		cart.RemoveZeroQuantityItems()

		cart.CalculateAmount()

		JSON(w, cart)
	}
}

func UpdateCartItemHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		form := new(CartForm)

		if err := Bind(form, r); err != nil {
			JSON(w, err, 400)
		}

		if err := form.Validate("PUT_ITEM"); err != nil {
			JSON(w, err, 422)
			return
		}

		cart, err := form.Cart(app)
		if err != nil {
			JSON(w, err, 422)
			return
		}
		if err := (CartRepository{app.DB}).SaveCart(cart); err != nil {
			JSON(w, err, 500)
			return
		}

		cart.RemoveZeroQuantityItems()

		cart.CalculateAmount()

		JSON(w, cart)
	}
}

func DeleteCartItemHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		form := new(CartForm)

		if err := Bind(form, r); err != nil {
			JSON(w, err, 400)
		}

		if err := form.Validate("DELETE"); err != nil {
			JSON(w, err, 422)
			return
		}

		cart, err := form.Cart(app)
		if err != nil {
			JSON(w, err, 422)
			return
		}
		if err := (CartRepository{app.DB}).DeleteCartItem(cart, c.ID()); err != nil {
			JSON(w, err, 500)
			return
		}

		cart.RemoveZeroQuantityItems()

		cart.CalculateAmount()

		JSON(w, cart)
	}
}
