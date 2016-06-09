package services_test

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/services"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"os"
)

var _ = Describe("TestImageService", func() {
	var _ = Describe("Test IsValidImage function", func() {
		It("is a valid jpeg image", func() {
			path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/golang.jpeg"
			file := ReturnMultipartFileheader(path)
			valid, err := (ImageService{}).IsValidImage(file)
			Expect(valid).To(Equal(true))
			Expect(err).To(BeNil())
		})

		It("return error because the image is too small", func() {
			path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/golang_small.png"
			file := ReturnMultipartFileheader(path)
			valid, err := (ImageService{}).IsValidImage(file)
			Expect(valid).To(Equal(false))
			Expect(err.Error()).To(ContainSubstring("Image size is too small, Width/Height's minimum value should be 500"))
		})

		It("return error because it is not a valid image", func() {
			path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/main.go"
			file := ReturnMultipartFileheader(path)
			valid, err := (ImageService{}).IsValidImage(file)
			Expect(valid).To(Equal(false))
			Expect(err.Error()).To(ContainSubstring("image: unknown format"))
		})
	})

	var _ = Describe("Test GetDetail function", func() {
		It("return valid detail image", func() {
			path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/golang.jpeg"
			outpath := "detail.png"

			file := ReturnMultipartFileheader(path)
			bytes := (ImageService{}).GetDetail(file, 550)

			ioutil.WriteFile(outpath, bytes, 0777)
			width, height := GetImageFileDimension(outpath)
			Expect(width).To(Equal(550))
			Expect(height).To(Equal(550))
			os.Remove(outpath)
		})

		It("return error bacause input file is invalid", func() {
			path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/product_factory.go"

			file := ReturnMultipartFileheader(path)
			bytes := (ImageService{}).GetDetail(file, 550)

			Expect(bytes).To(BeNil())
		})
	})

	var _ = Describe("Test GetThumbnail function", func() {
		It("return error bacause input file is invalid", func() {
			path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/product_factory.go"

			file := ReturnMultipartFileheader(path)
			bytes := (ImageService{}).GetThumbnail(file, 550)

			Expect(bytes).To(BeNil())
		})

		It("return valid thumbnail image", func() {
			path := os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/golang.jpeg"
			outpath := "thumbnail.png"

			file := ReturnMultipartFileheader(path)
			bytes := (ImageService{}).GetThumbnail(file, 320)

			ioutil.WriteFile(outpath, bytes, 0777)
			width, height := GetImageFileDimension(outpath)
			Expect(width).To(Equal(320))
			Expect(height).To(Equal(320))
			os.Remove(outpath)
		})
	})
})
