package handlers

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/form"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"

	"net/http"
)

func GetProductsDetailHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		// the purpose of form is to get form.Fields, so don't care about Binding errors
		form := new(GetProductForm)
		Bind(form, r)

		product, err := (ProductRepository{app.DB}).FindById(c.ID())
		if err != nil {
			JSON(w, err, 404)
			return
		}

		maps, err := FieldSelection(product, form.Fields)
		if err != nil {
			JSON(w, err)
			return
		}
		JSON(w, maps)
	}
}

func GetProductsHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		form := new(GetProductForm)

		if err := Bind(form, r); err != nil {
			form.Offset = 0
			form.Limit = 10
		}

		products, total, err := ProductRepository{app.DB}.GetPage(form.Offset, form.Limit, form.Search)
		if err != nil {
			JSON(w, err, 500)
			return
		}

		maps, err := FieldSelection(products.GetPaging(form.Offset, form.Limit, total), form.Fields)
		if err != nil {
			JSON(w, err)
			return
		}
		JSON(w, maps)
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

		maps, err := FieldSelection(product, form.Fields)
		if err != nil {
			JSON(w, err)
			return
		}
		JSON(w, maps, 201)
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

		maps, err := FieldSelection(product, form.Fields)
		if err != nil {
			JSON(w, err)
			return
		}
		JSON(w, maps, 200)
	}
}

func DeleteProductHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		product, err := (ProductRepository{app.DB}).DeleteProduct(c.ID())
		if err != nil {
			JSON(w, err, 500)
			return
		}
		JSON(w, product)
	}
}

func OptionHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		w.Header().Add("Access-Control-Allow-Origin", `*`)
		w.Header().Add("Access-Control-Allow-Methods", `GET, POST, PUT, DELETE, OPTIONS`)
		w.Header().Add("Access-Control-Allow-Headers", `content-type,accept`)
		w.Header().Add("Access-Control-Max-Age", "10")
		w.WriteHeader(http.StatusNoContent)
	}
}
