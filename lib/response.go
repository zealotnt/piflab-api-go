package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"image"
	"image/png"
	"net/http"
)

func JSON(w http.ResponseWriter, params ...interface{}) {
	setHTTPStatus(w, params)

	obj := params[0]

	switch obj.(type) {
	case error:
		obj = &Error{obj.(error)}
	}

	json.NewEncoder(w).Encode(obj)
}

func Image(w http.ResponseWriter, img image.Image) {
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		JSON(w, errors.New("unable to encode image."))
		return
	}

	// w.Header().Set("Content-Type", "image/png")
	if _, err := w.Write(buffer.Bytes()); err != nil {
		JSON(w, errors.New("unable to write image."))
	}
}

func setHTTPStatus(w http.ResponseWriter, params []interface{}) {
	if len(params) == 2 {
		status := params[1].(int)

		if status == 401 {
			w.Header().Set("WWW-Authenticate", `xBasic realm="fake"`)
		}

		w.WriteHeader(status)
	}
}
