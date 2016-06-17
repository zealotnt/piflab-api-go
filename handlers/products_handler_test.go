package handlers_test

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/handlers"
	"github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
)

type BindingTest struct {
	body   string
	expect string
}

type ValidateTest struct {
	body   string
	expect string
}

var _ = Describe("prduct_handlers Test", func() {
	GoodBucketName := os.Getenv("S3_BUCKET_NAME")
	BadBucketName := "wrong!!!"

	BeforeEach(func() {
		os.Setenv("S3_BUCKET_NAME", GoodBucketName)
	})

	AfterEach(func() {
		os.Setenv("S3_BUCKET_NAME", GoodBucketName)
	})

	var _ = Describe("GetProductsHandlerTest", func() {
		It("get products successfully, with status code 200", func() {
			response := Request("GET", "/products", "")
			Expect(response.Code).To(Equal(200))
		})

		It("gets products fail, because connection to db has been closed", func() {
			/* Close connection to database */
			app.Close()

			/* Fail to GET products */
			response := Request("GET", "/products", "")
			Expect(response.Code).To(Equal(500))
			Expect(response.Body).To(ContainSubstring("database is closed"))

			/* Connect again, others test cases still want database connection */
			app = lib.NewApp()
			app.AddRoutes(GetRoutes())
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
				{`{"name": "XBox","price": 70000,"provider": "Microsoft","rating": 6.0,"status": "sale"}`, `"Rating must be less than or equal to 5"`},
				{`{"name": "XBox","price": 70000,"provider": "Microsoft","rating": 3.5,"status": "on sale"}`, `"Status is invalid"`},
			}

			for _, test := range test_cases {
				response := Request("POST", "/products", test.body)
				Expect(response.Code).To(Equal(422))
				Expect(response.Body).To(ContainSubstring(test.expect))
			}
		})

		It("can't create a product, due to wrong AWS Bucket name (can't create image)", func() {
			os.Setenv("S3_BUCKET_NAME", BadBucketName)
			path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/golang.png"
			response := MultipartRequest("POST", "/products", extraParams, "image", path)
			Expect(response.Code).To(Equal(500))
			Expect(response.Body).To(ContainSubstring("NoSuchBucket: The specified bucket does not exist"))
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

		It("can't update a product, due to wrong AWS Bucket name (can't delete image)", func() {
			os.Setenv("S3_BUCKET_NAME", BadBucketName)
			path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/golang.png"
			response := MultipartRequest("PUT", getFirstAvailableUrl(), extraParams, "image", path)
			Expect(response.Code).To(Equal(500))
			Expect(response.Body).To(ContainSubstring("NoSuchBucket: The specified bucket does not exist"))
		})
	})

	var _ = Describe("DeleteProductHandlerTest", func() {
		It("delete fail, because no record invalid, and returns 500", func() {
			response := Request("DELETE", "/products/0", "")
			Expect(response.Code).To(Equal(500))
		})

		It("can't delete a product, due to wrong AWS Bucket name (can't delete image)", func() {
			os.Setenv("S3_BUCKET_NAME", BadBucketName)
			path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/golang.png"
			response := MultipartRequest("DELETE", getFirstAvailableUrl(), extraParams, "image", path)
			Expect(response.Code).To(Equal(500))
			Expect(response.Body).To(ContainSubstring("NoSuchBucket: The specified bucket does not exist"))
		})
	})
})
