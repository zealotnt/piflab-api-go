package handlers

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	// . "github.com/o0khoiclub0o/piflab-store-api-go/models/form"
	// . "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"

	"net/http"
)

func GetCartHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var cart Order

		// Forward it to carts service
		resp, body, err := RequestForwarder(r, app.CART_SERVICE, &cart)
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}
		if err != nil {
			JSON(w, err, resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, cart)
	}
}

func UpdateCartHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var cart Order

		// Forward it to carts service
		resp, body, err := RequestForwarder(r, app.CART_SERVICE, &cart)
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}
		if err != nil {
			JSON(w, err, resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, cart)
	}
}

func UpdateCartItemHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var cart Order

		// Forward it to carts service
		resp, body, err := RequestForwarder(r, app.CART_SERVICE, &cart)
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}
		if err != nil {
			JSON(w, err, resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, cart)
	}
}

func DeleteCartItemHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var cart Order

		// Forward it to carts service
		resp, body, err := RequestForwarder(r, app.CART_SERVICE, &cart)
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}
		if err != nil {
			JSON(w, err, resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, cart)
	}
}

func CheckoutCartHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var checkout CheckoutReturn

		// Forward it to carts service
		resp, body, err := RequestForwarder(r, app.CART_SERVICE, &checkout)
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}
		if err != nil {
			JSON(w, err, resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, checkout)
	}
}

func GetCheckoutHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var order_page OrderPage

		// Forward it to carts service
		resp, body, err := RequestForwarder(r, app.ORDER_SERVICE, &order_page)
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}
		if err != nil {
			JSON(w, err, resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, order_page)
	}
}

func GetCheckoutDetailHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var checkout_detail CheckoutReturn

		// Forward it to orders service
		resp, body, err := RequestForwarder(r, app.ORDER_SERVICE, &checkout_detail)
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}
		if err != nil {
			JSON(w, err, resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, checkout_detail)
	}
}

func UpdateCheckoutStatusHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var checkout_detail CheckoutReturn

		// Forward it to orders service
		resp, body, err := RequestForwarder(r, app.ORDER_SERVICE, &checkout_detail)
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}
		if err != nil {
			JSON(w, err, resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, checkout_detail)
	}
}
