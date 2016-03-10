package handlers

import (
	"net/http"
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"
)

func ProductsHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		products, err := ProductRepository{app.DB}.GetAll()
		
		if err != nil {
			JSON(w, err, 400)
			return
		} else {
			JSON(w, products)
		}
	
}}
