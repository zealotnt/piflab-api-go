package models_test

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"strconv"
)

type ProductSliceTest struct {
	description string
	url         string
	expect      string
}

var _ = Describe("Product Test", func() {
	var first bool
	var product *Product

	BeforeEach(func() {
		if first == false {
			first = true

			response := Request("GET", "/products", "")
			id := getFirstAvailableId(response)
			product, _ = (ProductRepository{app.DB}).FindById(id)
		}
	})

	var _ = Describe("GetImageUrlType Test", func() {
		It("returns valid image url, base on ORIGIN ImageType", func() {
			_, err := product.GetImageUrlType(ORIGIN)
			Expect(err).To(BeNil())
		})

		It("returns error when trying to get image url, because the param is out of scope of enum ImageType", func() {
			url, err := product.GetImageUrlType(99)
			Expect(url).To(Equal(""))
			Expect(err.Error()).To(ContainSubstring("field too short, minimum length 1: Key"))
		})
	})

	var _ = Describe("GetImagePath Test", func() {
		It("returns image url with image extension", func() {
			// rename Image file name to contain png extension
			product.Image = "dummyFile.png"
			path := product.GetImagePath(ORIGIN)
			expected := "products/" +
				strconv.FormatUint(uint64(product.Id), 10) +
				"/origin."
			Expect(path).To(ContainSubstring(expected))
			Expect(path).To(ContainSubstring(".png"))
		})

		It("returns image url without image extension", func() {
			// rename Image file name, so we don't use regex's result to give to file name
			product.Image = "dummyFile"
			path := product.GetImagePath(ORIGIN)
			expected := "products/" +
				strconv.FormatUint(uint64(product.Id), 10) +
				"/origin."
			Expect(path).To(ContainSubstring(expected))
			Expect(path).NotTo(ContainSubstring(".png"))
		})

		It("returns empty string, because the param is out of scope of enum ImageType", func() {
			url := product.GetImagePath(99)
			Expect(url).To(Equal(""))
		})
	})
})

var _ = Describe("ProductSlice Test", func() {
	It("test all cases of GetPaging-getPage", func() {
		product_counts, err := (ProductRepository{app.DB}).CountProduct()
		Expect(err).To(BeNil())

		test_cases := []ProductSliceTest{{
			/*1*/
			description: `Get only 1 product per page, with offset=0.` +
				`Return val should contain valid "next" field, and null "previous" field`,
			url:    `/products/offset=0&limit=1`,
			expect: `"paging":{"next":"/products/offset=1\u0026limit=1","previous":null}`}, {

			/*2*/
			description: `Get page with very big offset (bigger than maximum current product).
				The result should return null "next" field, and valid "previous" field, that can:
				+ Return all of the products, if "limit" > "product_counts"
				+ Return "limit" number of products, if "limit" < "products_counts"
				--> Return "limit" number of products, if "limit" < "products_counts"`,
			url: `/products/offset=` + strconv.FormatUint(uint64(product_counts), 10) +
				`&limit=1`,
			expect: `"paging":{"next":null,"previous":"/products/offset=` +
				strconv.FormatUint(uint64(product_counts-1), 10) +
				`\u0026limit=1"}`}, {

			/*3*/
			description: `Get page with very big offset (bigger than maximum current product).
				The result should return null "next" field, and valid "previous" field, that can:
				+ Return all of the products, if "limit" > "product_counts"
				+ Return "limit" number of products, if "limit" < "products_counts"
				--> Return all of the products, if "limit" > "product_counts"`,
			url: `/products/offset=` +
				strconv.FormatUint(uint64(product_counts), 10) +
				`&limit=` +
				strconv.FormatUint(uint64(product_counts), 10),
			expect: `"paging":{"next":null,"previous":"/products/offset=0\u0026limit=` +
				strconv.FormatUint(uint64(product_counts), 10)}, {

			/*4*/
			description: `Get page with offset in the "middle" position of products
				with the "limit" value that doesn't exceed the maximum "product_counts".
				So the result will contain both "next" field, and "previous" field`,
			url: `/products/offset=` + strconv.FormatUint(uint64(product_counts/2), 10) +
				`&limit=1`,
			expect: `"paging":{"next":"/products/offset=` +
				strconv.FormatUint(uint64(product_counts/2+1), 10) + `\u0026limit=1",` +
				`"previous":"/products/offset=` +
				strconv.FormatUint(uint64(product_counts/2-1), 10) + `\u0026limit=1`},
		}

		for _, test := range test_cases {
			By(test.description)
			response := Request("GET", test.url, "")
			body, _ := ioutil.ReadAll(response.Body)
			Expect(response.Code).To(Equal(200))
			Expect(body).To(ContainSubstring(test.expect))
		}
	})
})
