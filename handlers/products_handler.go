package handlers

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/form"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"
	"net/http"
)

func GetProductsHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		products, err := ProductRepository{app.DB}.GetAll()

		if err != nil {
			JSON(w, err, 500)
			return
		}

		JSON(w, products)
	}
}

func CreateProductHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		form := new(CreateProductForm)

		if err := Bind(form, r); err != nil {
			JSON(w, err, 400)
			return
		}

		if err := form.Validate(); err != nil {
			JSON(w, err, 422)
			return
		}

		product := form.Product()
		if err := (ProductRepository{app.DB}).SaveProduct(product); err != nil {
			JSON(w, err, 500)
			return
		}
		JSON(w, product, 201)
	}
}

func UpdateProductHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		product, err := (ProductRepository{app.DB}).FindById(c.ID())
		if err != nil {
			JSON(w, err, 404)
			return
		}

		form := new(UpdateProductForm)

		if err := Bind(form, r); err != nil {
			JSON(w, err, 400)
			return
		}

		if err := form.Validate(); err != nil {
			JSON(w, err, 422)
			return
		}

		form.Assign(product)
		if err := (ProductRepository{app.DB}).SaveProduct(product); err != nil {
			JSON(w, err, 500)
			return
		}
		JSON(w, product)
	}
}
