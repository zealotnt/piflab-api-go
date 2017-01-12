package handlers

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"

	"net/http"
)

func ProductForwardHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		// Forward it to service
		resp, body, err := RequestForwarder(r, app.PRODUCT_SERVICE, nil)
		if err != nil {
			if resp == nil {
				JSON(w, err)
				return
			}
			JSON(w, err, resp.StatusCode)
			return
		}
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}

		// Temporary not support field selection
		WriteBody(w, body)
	}
}

func CreateProductHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		// Forward it to service
		resp, body, err := RequestForwarder(r, app.PRODUCT_SERVICE, nil)
		if err != nil {
			if resp == nil {
				JSON(w, err)
				return
			}
			JSON(w, err, resp.StatusCode)
			return
		}
		if resp.Status != "201 Created" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}

		// Temporary not support field selection
		WriteBody(w, body)
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
