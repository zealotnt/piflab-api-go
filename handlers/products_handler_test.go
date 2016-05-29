package handlers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type BindingTest struct {
	body   string
	expect string
}

type ValidateTest struct {
	body   string
	expect string
}

var _ = Describe("GetProductsHandlerTest", func() {
	It("returns 200", func() {
		response := Request("GET", "/products", "")

		Expect(response.Code).To(Equal(200))
	})
})

var _ = Describe("CreateProductHandlerTest", func() {
	It("has erroneous binding result, and returns 400", func() {
		var test_cases = []BindingTest{
			{`{"name": "XBox","price": "70000","provider": "Microsoft","rating": 3.5,"status": "sale"}`, `"json: cannot unmarshal string into Go value of type int"`},
			{`{"name": "XBox","price": 70000,"provider": "Microsoft","rating": "3.5","status": "sale"}`, `"json: cannot unmarshal string into Go value of type float32"`},
		}

		for _, test := range test_cases {
			response := Request("POST", "/products", test.body)
			Expect(response.Code).To(Equal(400))
			Expect(response.Body).To(ContainSubstring(test.expect))
		}
	})

	It("has erroneous validation result, and returns 422", func() {
		var test_cases = []ValidateTest{
			{`{"price": 70000,"provider": "Microsoft","rating": 3.5,"status": "sale"}`, `"Name is required"`},
			{`{"name": "XBox","provider": "Microsoft","rating": 3.5,"status": "sale"}`, `"Price is required"`},
			{`{"name": "XBox","price": 70000,"rating": 3.5,"status": "sale"}`, `"Provider is required"`},
			{`{"name": "XBox","price": 70000,"provider": "Microsoft","status": "sale"}`, `"Rating is required"`},
			{`{"name": "XBox","price": 70000,"provider": "Microsoft","rating": 3.5}`, `"Status is required"`},
			{`{"name": "XBox","price": 70000,"provider": "Microsoft","rating": 3.5,"status": "sale"}`, `"Image is required"`},
			{`{"name": "XBox","price": 70000,"provider": "Microsoft","rating": 6.0,"status": "sale"}`, `"Rating must be less than or equal to 5"`},
			{`{"name": "XBox","price": 70000,"provider": "Microsoft","rating": 3.5,"status": "on sale"}`, `"Status is invalid"`},
		}

		for _, test := range test_cases {
			response := Request("POST", "/products", test.body)
			Expect(response.Code).To(Equal(422))
			Expect(response.Body).To(ContainSubstring(test.expect))
		}
	})
})

var _ = Describe("UpdateProductHandlerTest", func() {
	It("has invalid product id request, and returns 404", func() {
		response := Request("PUT", "/products/abc", `{"name": "XBox"}`)

		Expect(response.Code).To(Equal(404))
		Expect(response.Body).To(ContainSubstring(`"record not found"`))
	})

	It("has product id zero request, and returns 404", func() {
		response := Request("PUT", "/products/0", `{"name": "XBox"}`)

		Expect(response.Code).To(Equal(404))
		Expect(response.Body).To(ContainSubstring(`"record not found"`))
	})

	It("has erroneous binding result, and returns 400", func() {
		var test_cases = []BindingTest{
			{`{"rating": "3.4"}`, `"json: cannot unmarshal string into Go value of type float32"`},
			{`{"price": "123"}`, `"json: cannot unmarshal string into Go value of type int"`},
		}

		for _, test := range test_cases {
			response := Request("PUT", getFirstAvailableUrl(), test.body)
			Expect(response.Code).To(Equal(400))
			Expect(response.Body).To(ContainSubstring(test.expect))
		}
	})

	It("has erroneous validation result, and returns 422", func() {
		var test_cases = []ValidateTest{
			{`{"rating": 5.1}`, `"Rating must be less than or equal to 5"`},
			{`{"status": "on sale"}`, `"Status is invalid"`},
		}

		for _, test := range test_cases {
			response := Request("PUT", getFirstAvailableUrl(), test.body)
			Expect(response.Code).To(Equal(422))
			Expect(response.Body).To(ContainSubstring(test.expect))
		}
	})

	It("updates success, and returns 200", func() {
		response := Request("PUT", getFirstAvailableUrl(), `{"rating": 4.0}`)
		Expect(response.Code).To(Equal(200))
	})

})
