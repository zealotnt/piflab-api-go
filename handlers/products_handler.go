package handlers

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"
	"net/http"

	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
)

func GetAllProductsHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		products, err := ProductRepository{app.DB}.GetAll()

		if err != nil {
			JSON(w, err, 400)
			return
		}

		JSON(w, products)
	}
}

func CreateProductHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {

		form := new(ProductForm)
		if err := Bind(form, r); err != nil {
			JSON(w, err, 400)
			return
		}

		if err := form.Validate(); err != nil {
			JSON(w, err, 422)
			return
		}

		product := form.Product()

		err := ProductRepository{app.DB}.CreateProduct(&product)

		if err != nil {
			JSON(w, err, 422)
			return
		}

		JSON(w, product, 201)
	}
}
