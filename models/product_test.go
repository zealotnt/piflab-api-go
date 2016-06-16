package models_test

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Product", func() {
	It("return image url", func() {
		response := Request("GET", "/products", "")
		id := getFirstAvailableId(response)
		product, _ := (ProductRepository{app.DB}).FindById(id)
		product.GetImageUrlType(ORIGIN)

		url, err := product.GetImageUrlType(99)
		Expect(url).To(Equal(""))
		Expect(err.Error()).To(ContainSubstring("field too short, minimum length 1: Key"))
	})
})

var _ = Describe("ProductWithImgExtension", func() {
	It("return image url with image extension", func() {
		response := Request("GET", "/products", "")
		id := getFirstAvailableId(response)
		product, _ := (ProductRepository{app.DB}).FindById(id)
		// rename Image file name, so we don't use regex's result to give to file name
		product.Image = "dummyFile"
		product.GetImageUrlType(ORIGIN)
	})
})
