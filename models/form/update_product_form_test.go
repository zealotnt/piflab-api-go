package models_test

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/form"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"os"
	"strconv"
)

var _ = Describe("UpdateProductFormFieldMap", func() {
	It("dummy testing", func() {
		dummy := new(http.Request)
		form := new(UpdateProductForm)
		form.FieldMap(dummy)
	})
})

var _ = Describe("ValidateUpdateProductForm", func() {
	var extraParams = map[string]string{}
	var form = UpdateProductForm{}

	BeforeEach(func() {
		form = UpdateProductForm{}
		extraParams = map[string]string{
			"name":     name,
			"price":    strconv.FormatInt(int64(price), 10),
			"provider": provider,
			"rating":   strconv.FormatFloat(float64(rating), 'f', 1, 32),
			"status":   status,
		}
	})

	It("has exceeded rating limit", func() {
		extraParams["rating"] = strconv.FormatFloat(float64(ratingBig), 'f', 1, 32)
		BindForm(&form, extraParams, "")
		err := form.Validate()
		Expect(err.Error()).To(ContainSubstring("Rating must be less than or equal to 5"))
	})

	It("has invalid status", func() {
		extraParams["status"] = invalidStatus
		BindForm(&form, extraParams, "")
		err := form.Validate()
		Expect(err.Error()).To(ContainSubstring("Status is invalid"))
	})

	It("has invalid image extension", func() {
		path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/main.go"
		err := BindForm(&form, extraParams, path)
		Expect(err).To(BeNil())

		err = form.Validate()
		Expect(err.Error()).To(ContainSubstring("Image extension is invalid"))
	})

	It("updates successfully", func() {
		path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/golang.jpeg"
		err := BindForm(&form, extraParams, path)
		Expect(err).To(BeNil())

		err = form.Validate()
		Expect(err).To(BeNil())
	})

	It("returns a Product from Form using Assign", func() {
		path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/golang.jpeg"
		err := BindForm(&form, extraParams, path)
		Expect(err).To(BeNil())

		err = form.Validate()
		Expect(err).To(BeNil())

		product := new(Product)
		form.Assign(product)
		Expect(product.Name).To(Equal(name))
		Expect(product.Price).To(Equal(price))
		Expect(product.Provider).To(Equal(provider))
		Expect(product.Rating).To(Equal(rating))
		Expect(product.Status).To(Equal(status))
		Expect(len(product.ImageData)).To(Equal(len(form.ImageData())))
	})

})
