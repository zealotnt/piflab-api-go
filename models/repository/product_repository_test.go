package repository_test

import (
	// "github.com/icrowley/fake"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
)

var _ = Describe("ProductRepositoryTest", func() {
	testName := "testProduct"
	testNewName := "testNewProductName"
	testPrice := 1000
	testProvider := "testProvider"
	testRating := float32(3.0)
	testStatus := "sale"
	testDetail := "Lorem Ipsum"
	testImageData := []byte("Some miscellaneous data")
	testImageThumbnail := []byte("Some miscellaneous thumbnail data")
	testImageDetail := []byte("Some miscellaneous detail data")

	GoodBucketName := os.Getenv("S3_BUCKET_NAME")
	BadBucketName := "wrong!!!"

	var product Product
	BeforeEach(func() {
		product = Product{
			Name:               testName,
			Price:              testPrice,
			Provider:           testProvider,
			Rating:             testRating,
			Status:             testStatus,
			Detail:             testDetail,
			ImageData:          testImageData,
			ImageThumbnailData: testImageThumbnail,
			ImageDetailData:    testImageDetail,
		}
	})

	It("Test all", func() {
		os.Setenv("S3_BUCKET_NAME", GoodBucketName)

		/* Create fail due to over limit of characters */
		// product.Id = 0
		// product.Name = fake.WordsN(300)
		// err := ProductRepository{app.DB}.SaveProduct(&product)
		// Expect(err.Error()).NotTo(BeNil())

		/* Then create ok now */
		product.Name = testName
		err := ProductRepository{app.DB}.SaveProduct(&product)
		Expect(err).To(BeNil())

		/* Get previous created product */
		temp_product, err := ProductRepository{app.DB}.FindById(product.Id)
		Expect(err).To(BeNil())
		Expect(temp_product.Id).To(Equal(product.Id))

		/* Update product with new name */
		product.Name = testNewName
		err = ProductRepository{app.DB}.SaveProduct(&product)
		Expect(err).To(BeNil())
		Expect(product.Name).To(Equal(testNewName))

		/* Try to update product with exceed length name */
		// product.Name = fake.WordsN(300)
		// err = ProductRepository{app.DB}.SaveProduct(&product)
		// Expect(err).NotTo(BeNil())

		/* Make sure the exceed name won't updated, but still has the old name */
		temp_product, err = ProductRepository{app.DB}.FindById(product.Id)
		Expect(err).To(BeNil())
		Expect(temp_product.Name).To(Equal(testNewName))

		/* Delete product */
		_, err = ProductRepository{app.DB}.DeleteProduct(product.Id)
		Expect(err).To(BeNil())

		/* Find it again, of course it fails  */
		temp_product, err = ProductRepository{app.DB}.FindById(product.Id)
		Expect(err.Error()).To(ContainSubstring("record not found"))
		Expect(temp_product).To(BeNil())

		/* Try to delete it, of course it fails */
		_, err = ProductRepository{app.DB}.DeleteProduct(product.Id)
		Expect(err.Error()).To(ContainSubstring("record not found"))
	})

	It("Fail to create/update/delete due to wrong AWS key", func() {
		/* Create a temporary record */
		err := ProductRepository{app.DB}.SaveProduct(&product)
		Expect(err).To(BeNil())

		os.Setenv("S3_BUCKET_NAME", BadBucketName)

		/* Fail to create */
		temp_product := product
		temp_product.Id = 0
		err = ProductRepository{app.DB}.SaveProduct(&temp_product)
		Expect(err.Error()).To(ContainSubstring("NoSuchBucket: The specified bucket does not exist"))

		/* Fail to update */
		err = ProductRepository{app.DB}.SaveProduct(&product)
		Expect(err.Error()).To(ContainSubstring("NoSuchBucket: The specified bucket does not exist"))

		/* Fail to delete */
		_, err = ProductRepository{app.DB}.DeleteProduct(product.Id)
		Expect(err.Error()).To(ContainSubstring("NoSuchBucket: The specified bucket does not exist"))

		/* Teardown temporary record */
		os.Setenv("S3_BUCKET_NAME", GoodBucketName)
		_, err = ProductRepository{app.DB}.DeleteProduct(product.Id)
		Expect(err).To(BeNil())
	})

	Describe("Test GetAll and CountProduct", func() {
		It("return same number of element", func() {
			product, err := ProductRepository{app.DB}.GetAll()
			Expect(err).To(BeNil())

			count, err := ProductRepository{app.DB}.CountProduct()
			Expect(err).To(BeNil())

			Expect(len(*product)).To(Equal(count))
		})
	})

})
