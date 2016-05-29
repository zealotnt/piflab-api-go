package models_test

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/form"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
)

var _ = Describe("ProductTest", func() {
	It("TestGetDataFunc", func() {
		form := new(ProductForm)
		path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/golang.jpeg"
		err := BindForm(form, nil, path)
		Expect(err).To(BeNil())
		Expect(len(form.ImageData())).To(Equal(getFileSize(path)))
	})
})
