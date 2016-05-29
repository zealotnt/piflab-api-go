package models_test

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/form"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"os"
	"strconv"
)

var _ = Describe("CreateProductFormFieldMap", func() {
	It("requires name", func() {
		dummy := new(http.Request)
		form := new(CreateProductForm)
		form.FieldMap(dummy)
	})
})

var _ = Describe("ValidateCreateProductForm", func() {
	var extraParams = map[string]string{}
	var form = CreateProductForm{}

	BeforeEach(func() {
		form = CreateProductForm{}
		extraParams = map[string]string{
			"name":     name,
			"price":    strconv.FormatInt(int64(price), 10),
			"provider": provider,
			"rating":   strconv.FormatFloat(float64(rating), 'f', 1, 32),
			"status":   status,
		}
	})

	It("requires name", func() {
		delete(extraParams, "name")
		BindForm(&form, extraParams, "")
		err := form.Validate()
		Expect(err.Error()).To(ContainSubstring("Name is required"))
	})

	It("requires price", func() {
		delete(extraParams, "price")
		BindForm(&form, extraParams, "")
		err := form.Validate()
		Expect(err.Error()).To(ContainSubstring("Price is required"))
	})

	It("requires provider", func() {
		delete(extraParams, "provider")
		BindForm(&form, extraParams, "")
		err := form.Validate()
		Expect(err.Error()).To(ContainSubstring("Provider is required"))
	})

	It("requires rating", func() {
		delete(extraParams, "rating")
		BindForm(&form, extraParams, "")
		err := form.Validate()
		Expect(err.Error()).To(ContainSubstring("Rating is required"))
	})

	It("has exceeded the rating limit", func() {
		extraParams["rating"] = strconv.FormatFloat(float64(ratingBig), 'f', 1, 32)
		BindForm(&form, extraParams, "")
		err := form.Validate()
		Expect(err.Error()).To(ContainSubstring("Rating must be less than or equal to 5"))
	})

	It("requires status", func() {
		delete(extraParams, "status")
		BindForm(&form, extraParams, "")
		err := form.Validate()
		Expect(err.Error()).To(ContainSubstring("Status is required"))
	})

	It("has invalid status", func() {
		extraParams["status"] = invalidStatus
		BindForm(&form, extraParams, "")
		err := form.Validate()
		Expect(err.Error()).To(ContainSubstring("Status is invalid"))
	})

	It("requires image", func() {
		BindForm(&form, extraParams, "")
		err := form.Validate()
		Expect(err.Error()).To(ContainSubstring("Image is required"))
	})

	It("has invalid image extension", func() {
		path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/main.go"
		err := BindForm(&form, extraParams, path)
		Expect(err).To(BeNil())

		err = form.Validate()
		Expect(err.Error()).To(ContainSubstring("Image extension is invalid"))
	})

	It("is success", func() {
		path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/golang.jpeg"
		err := BindForm(&form, extraParams, path)
		Expect(err).To(BeNil())

		err = form.Validate()
		Expect(err).To(BeNil())

		product := form.Product()
		Expect(product).NotTo(BeNil())
	})
})
