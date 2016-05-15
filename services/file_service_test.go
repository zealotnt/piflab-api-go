package services_test

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/services"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"net/http"
)

var _ = Describe("FileSeriveTestSuccessFlow", func() {
	testKey := "testKey"
	testData := []byte("Some miscellaneous data")
	service := FileService{}

	It("Do all test", func() {
		err := service.SaveFile(testData, testKey)
		Expect(err).To(BeNil())

		data, err := service.GetFile(testKey)
		Expect(err).To(BeNil())
		Expect(data).To(Equal(testData))

		url, err := service.GetProtectedUrl(testKey, 15)
		Expect(err).To(BeNil())

		response, err := http.Get(url)
		defer response.Body.Close()
		Expect(err).To(BeNil())
		contents, err := ioutil.ReadAll(response.Body)
		Expect(err).To(BeNil())
		Expect(testData).To(Equal(contents))

		err = service.DeleteFile(testKey)
		Expect(err).To(BeNil())
	})

	It("Will fail due to invalid param", func() {
		err := service.SaveFile(nil, testKey)
		Expect(err.Error()).To(ContainSubstring("File content is required"))

		err = service.SaveFile(testData, "")
		Expect(err.Error()).To(ContainSubstring("Key is required"))
	})
})
