package handlers_test

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/handlers"
	"github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"
	"io/ioutil"
	"os"
)

type ParamBindingTest struct {
	param  string
	expect string
}

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

	var _ = Describe("GetAllProductsHandler Test", func() {
		It("get products successfully, with status code 200", func() {
			response := Request("GET", "/products", "")
			Expect(response.Code).To(Equal(200))
			Expect(response.Header().Get(`Content-Type`)).To(Equal(`application/json`))
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

	var _ = Describe("GetPageProductsHandler Test", func() {
		It("has erroneous parameters binding result, returns 400", func() {
			var test_cases = []ParamBindingTest{
				{`/products/offset=0&limit=0`, `Limit must bigger than 0`},
				{`/products/offset=0&limit=-1`, `Limit must bigger than 0`},
				{`/products/offset=0&limit=-1ab`, `Error when parsing limit parameter`},
				{`/products/offset=-1&limit=0`, `Offset must bigger than or equal to 0`},
				{`/products/offset=-1a&limit=0`, `Error when parsing offset parameter`},
			}

			for _, test := range test_cases {
				response := Request("GET", test.param, "")
				Expect(response.Code).To(Equal(400))
				Expect(response.Body).To(ContainSubstring(test.expect))
			}
		})

		It("gets products fail, because connection to db has been closed", func() {
			/* Close connection to database */
			app.Close()

			/* Fail to GET products */
			response := Request("GET", `/products/offset=0&limit=1`, "")
			Expect(response.Code).To(Equal(500))
			Expect(response.Body).To(ContainSubstring("database is closed"))

			/* Connect again, others test cases still want database connection */
			app = lib.NewApp()
			app.AddRoutes(GetRoutes())
		})

		It("gets first product successfully", func() {
			/* Get a product */
			response := Request("GET", `/products/offset=0&limit=1`, "")
			Expect(response.Code).To(Equal(200))
			Expect(response.Header().Get(`Content-Type`)).To(Equal(`application/json`))

			/* Parse response's body */
			body, _ := ioutil.ReadAll(response.Body)

			/* Deserialize json */
			product_page := ProductPage{}
			err := json.Unmarshal(body, &product_page)

			/* Len should be equal to 1, and no error */
			Expect(len(*product_page.Data)).To(Equal(1))
			Expect(err).To(BeNil())
		})
	})

	var _ = Describe("CreateProductHandler Test", func() {
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

	var _ = Describe("UpdateProductHandler Test", func() {
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
			Expect(response.Header().Get(`Content-Type`)).To(Equal(`application/json`))
		})

		It("can't update a product, due to wrong AWS Bucket name (can't delete image)", func() {
			os.Setenv("S3_BUCKET_NAME", BadBucketName)
			path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/golang.png"
			response := MultipartRequest("PUT", getFirstAvailableUrl(), extraParams, "image", path)
			Expect(response.Code).To(Equal(500))
			Expect(response.Body).To(ContainSubstring("NoSuchBucket: The specified bucket does not exist"))
		})
	})

	var _ = Describe("DeleteProductHandler Test", func() {
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

	var _ = Describe("OptionHandler Test", func() {
		It("returns with neccessary headers", func() {
			response := Request("OPTIONS", "/", "")
			Expect(response.Code).To(Equal(204))
			Expect(response.Header().Get("Access-Control-Allow-Origin")).To(Equal(`*`))
			Expect(response.Header().Get("Access-Control-Allow-Methods")).To(Equal(`GET, POST, PUT, DELETE, OPTIONS`))
			Expect(response.Header().Get("Access-Control-Allow-Headers")).To(Equal(`content-type,accept`))
			Expect(response.Header().Get("Access-Control-Max-Age")).To(Equal("10"))
		})

		It("returns with any format of request url", func() {
			urls_test := []string{
				`/`,
				`/123`,
				`/abc`,
				`/abc/123`,
				`/some+space`,
				`/some%20space`,
			}
			for _, url := range urls_test {
				response := Request("OPTIONS", url, "")
				Expect(response.Code).To(Equal(204))
				Expect(response.Header().Get("Access-Control-Allow-Origin")).To(Equal(`*`))
				Expect(response.Header().Get("Access-Control-Allow-Methods")).To(Equal(`GET, POST, PUT, DELETE, OPTIONS`))
				Expect(response.Header().Get("Access-Control-Allow-Headers")).To(Equal(`content-type,accept`))
				Expect(response.Header().Get("Access-Control-Max-Age")).To(Equal("10"))
			}
		})
	})
})
