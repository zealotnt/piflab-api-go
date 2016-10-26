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

		if err := form.Validate("GET", app); err != nil {
			JSON(w, err, 401)
			return
		}

		order, err := (OrderRepository{app}).GetOrder(*form.AccessToken)
		if err != nil {
			JSON(w, err, 500)
			return
		}
		order.EraseAccessToken()

		maps, err := FieldSelection(order, form.Fields)
		if err != nil {
			JSON(w, err)
			return
		}
		JSON(w, maps)
	}
}

func UpdateCartHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		form := new(CartForm)

		if err := Bind(form, r); err != nil {
			JSON(w, err, 400)
		}

		if err := form.Validate("PUT_CART", app); err != nil {
			JSON(w, err, 422)
			return
		}

		order, err := form.Order(app)
		if err != nil {
			JSON(w, err, 424)
			return
		}
		if err := (OrderRepository{app}).SaveOrder(order); err != nil {
			JSON(w, err, 500)
			return
		}

		order.RemoveZeroQuantityItems()

		maps, err := FieldSelection(order, form.Fields)
		if err != nil {
			JSON(w, err)
			return
		}
		JSON(w, maps)
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

		order, err := form.Order(app, c.ID())
		if err != nil {
			JSON(w, err, 424)
			return
		}
		if err := (OrderRepository{app}).SaveOrder(order); err != nil {
			JSON(w, err, 500)
			return
		}

		order.RemoveZeroQuantityItems()

		maps, err := FieldSelection(order, form.Fields)
		if err != nil {
			JSON(w, err)
			return
		}
		JSON(w, maps)
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

		order, err := form.Order(app)
		if err != nil {
			JSON(w, err, 424)
			return
		}
		if err := (OrderRepository{app}).DeleteOrderItem(order, c.ID()); err != nil {
			JSON(w, err, 500)
			return
		}

		order.RemoveZeroQuantityItems()

		maps, err := FieldSelection(order, form.Fields)
		if err != nil {
			JSON(w, err)
			return
		}
		JSON(w, maps)
	}
}

func CheckoutCartHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		form := new(CheckoutCartForm)

		if err := Bind(form, r); err != nil {
			JSON(w, err, 400)
			return
		}

		if err := form.Validate(); err != nil {
			JSON(w, err, 422)
			return
		}

		order, err := form.Order(app)
		if err != nil {
			JSON(w, err, 424)
			return
		}

		if err := (OrderRepository{app}).CheckoutOrder(order); err != nil {
			JSON(w, err, 500)
			return
		}

		ret := order.ReturnCheckoutRequest()

		maps, err := FieldSelection(ret, form.Fields)
		if err != nil {
			JSON(w, err)
			return
		}
		JSON(w, maps)
	}
}

func GetCheckoutHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		form := new(GetCheckoutForm)

		if err := Bind(form, r); err != nil {
			form.Offset = 0
			form.Limit = 10
		}

		if err := form.Validate(); err != nil {
			JSON(w, err, 422)
			return
		}

		orders_by_pages, err := OrderRepository{app}.GetPage(form.Offset, form.Limit, *form.Status, form.SortField, form.SortOrder, form.Search)
		if err != nil {
			JSON(w, err, 500)
			return
		}
		JSON(w, orders_by_pages)
	}
}

func GetCheckoutDetailHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		// the purpose of form is to get form.Fields, so don't care about Binding errors
		form := new(GetCheckoutForm)
		Bind(form, r)

		order, err := (OrderRepository{app}).FindByOrderId(c.Params["id"])
		if err != nil {
			JSON(w, err, 404)
			return
		}

		ret := order.ReturnCheckoutRequest()

		maps, err := FieldSelection(ret, form.Fields)
		if err != nil {
			JSON(w, err)
			return
		}
		JSON(w, maps)
	}
}

func UpdateCheckoutStatusHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		form := new(UpdateCheckoutForm)
		if err := Bind(form, r); err != nil {
			JSON(w, err, 400)
			return
		}

		if err := form.Validate(); err != nil {
			JSON(w, err, 422)
			return
		}

		order, err := form.Order(app, c.Params["id"])
		if err != nil {
			JSON(w, err, 424)
			return
		}

		if err := (OrderRepository{app}).SaveOrder(order); err != nil {
			JSON(w, err, 500)
			return
		}

		ret := order.ReturnCheckoutRequest()

		maps, err := FieldSelection(ret, form.Fields)
		if err != nil {
			JSON(w, err)
			return
		}
		JSON(w, maps)
	}
}
