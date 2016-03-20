package lib

import (
	"net/http"
	"github.com/mholt/binding"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
)

func Bind(form *ProductForm, r *http.Request) error {
	errs := binding.Bind(r, form)

	if errs.Len() > 0 {
		return errs
	}

	return nil
}
