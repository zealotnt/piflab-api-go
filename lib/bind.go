package lib

import (
	"github.com/mholt/binding"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	"net/http"
)

func Bind(form *ProductForm, r *http.Request) error {
	errs := binding.Bind(r, form)

	if errs.Len() > 0 {
		return errs
	}

	return nil
}
