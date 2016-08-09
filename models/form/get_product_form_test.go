package models_test

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/form"
	. "github.com/onsi/ginkgo"

	"net/http"
)

var _ = Describe("GetProductFormFieldMap", func() {
	It("just for coverage", func() {
		dummy := new(http.Request)
		form := new(GetProductForm)
		form.FieldMap(dummy)
	})
})
