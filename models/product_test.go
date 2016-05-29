package models_test

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Product", func() {
	It("return image url", func() {
		response := Request("GET", "/products", "")
		id := getFirstAvailableId(response)
		product, _ := (ProductRepository{app.DB}).FindById(id)
		product.GetImageUrl()
	})
})

var _ = Describe("ProductWithImgExtension", func() {
	It("return image url with image extension", func() {
		response := Request("GET", "/products", "")
		id := getFirstAvailableId(response)
		product, _ := (ProductRepository{app.DB}).FindById(id)
		product.Image = "dummyFile"
		product.GetImageUrl()
	})
})
