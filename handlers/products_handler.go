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

		product, err := (ProductRepository{app}).FindById(c.ID())
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

		products_by_pages, err := ProductRepository{app}.GetPage(form.Offset, form.Limit, form.Search)
		if err != nil {
			JSON(w, err, 500)
			return
		}

		// Get the fully maps
		maps, err := FieldSelection(products_by_pages, "")
		if err != nil {
			JSON(w, err, 503)
			return
		}

		// If the product list is nil, return right away
		if products_by_pages.Data == nil {
			JSON(w, maps)
			return
		}

		// Filter the "data"'s fields
		var data_maps []map[string]interface{}
		for idx, _ := range *products_by_pages.Data {
			var data_in_map map[string]interface{}
			data := (*products_by_pages.Data)[idx]
			data_in_map, err = FieldSelection(data, form.Fields)
			if err != nil {
				JSON(w, err, 503)
				return
			}
			data_maps = append(data_maps, data_in_map)
		}
		// Give the filtered data to the output
		maps["data"] = data_maps
		JSON(w, maps)
	}
}

func CreateProductHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		resp, err := RequestForwarder(r, "products:9901")
		if err != nil {
			JSON(w, err, 500)
		}
		JSON(w, resp.Body)
		// PR_DUMP(r)
		return
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
		if err := (ProductRepository{app}).SaveProduct(product); err != nil {
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
		product, err := (ProductRepository{app}).FindById(c.ID())
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
		if err := (ProductRepository{app}).SaveProduct(product); err != nil {
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
		product, err := (ProductRepository{app}).DeleteProduct(c.ID())
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
