package lib

import (
	"github.com/mholt/binding"
	"net/http"
)

func Bind(form binding.FieldMapper, r *http.Request) error {
	errs := binding.Bind(r, form)

	if errs.Len() > 0 {
		return errs
	}

	return nil
}
