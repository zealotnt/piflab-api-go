package models_test

import (
	"github.com/mholt/binding"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

var name = "xbox"
var price = 70000
var provider = "Microsoft"
var rating = float32(3.5)
var ratingBig = float32(5.1)
var ratingLessThanZero = float32(-0.5)
var status = "sale"
var invalidStatus = "on sale"
var detail = "some text"
var image = new(multipart.FileHeader)

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Form Suite")
}

var _ = BeforeSuite(func() {
})

var _ = AfterSuite(func() {
})

func getFileSize(path string) int {
	file, err := os.Open(path)
	if err != nil {
		return 0
	}
	fi, err := file.Stat()
	if err != nil {
		return 0
	}
	return int(fi.Size())
}

func createHttpRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	file, err := os.Open(path)
	if err == nil {
		part, _ := writer.CreateFormFile(paramName, filepath.Base(path))
		io.Copy(part, file)
	}
	defer file.Close()

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	contentType := writer.FormDataContentType()
	request.Header.Set("Content-Type", contentType)
	return request, err
}

func BindForm(form binding.FieldMapper, params map[string]string, image_path string) error {
	request, err := createHttpRequest("", params, "image", image_path)
	if err != nil {
		return err
	}

	return binding.Bind(request, form)
}
