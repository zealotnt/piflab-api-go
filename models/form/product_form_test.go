package models_test

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/form"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
	"strconv"
)

var _ = Describe("ProductTest", func() {
	var extraParams = map[string]string{}

	BeforeEach(func() {
		extraParams = map[string]string{
			"name":     name,
			"price":    strconv.FormatInt(int64(price), 10),
			"provider": provider,
			"rating":   strconv.FormatFloat(float64(rating), 'f', 1, 32),
			"status":   status,
			"detail":   detail,
		}
	})

	var _ = Describe("Tests ImageData function", func() {
		It("return nil, because there's no image field in form", func() {
			form := new(ProductForm)
			err := BindForm(form, nil, "")
			Expect(err).To(BeNil())
			Expect(form.ImageData()).To(BeNil())
		})

		It("returns sucessfully", func() {
			form := new(ProductForm)
			path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/golang.jpeg"
			err := BindForm(form, nil, path)
			Expect(err).To(BeNil())
			Expect(len(form.ImageData())).To(Equal(getFileSize(path)))
		})
	})

	It("TestGetDataFunc without Image, so the image relating fields are nil", func() {
		form := new(ProductForm)
		err := BindForm(form, extraParams, "")
		Expect(err).To(BeNil())
		product := form.Product()
		Expect(product.ImageUrl).To(BeNil())
		Expect(product.ImageThumbnailUrl).To(BeNil())
		Expect(product.ImageDetailUrl).To(BeNil())
	})
})
