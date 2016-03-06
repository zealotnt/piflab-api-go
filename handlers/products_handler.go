package handlers

import (
	"encoding/json"
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	"net/http"
)

func ProductsHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		products := [2]Product{{
			Price:     70000,
			Name:      "XBox",
			Provider:  "Microsoft",
			Rating:    3,
			Status:    "sale",
			ImageUrl:  "img/catalog/1.png",
			DetailUrl: "/product/1",
		}, {
			Price:     50000,
			Name:      "PS3",
			Provider:  "Sony",
			Rating:    4,
			Status:    "sale",
			ImageUrl:  "img/catalog/2.png",
			DetailUrl: "/product/2",
		},
		}

		json.NewEncoder(w).Encode(products)
	}
}
