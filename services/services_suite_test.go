package services_test

import (
	"github.com/mholt/binding"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/form"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"image"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Services Suite")
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

func ReturnMultipartFileheader(image_path string) *multipart.FileHeader {
	var extraParams = map[string]string{
		"name":     "name",
		"price":    "123",
		"provider": "mic",
		"rating":   "3.5",
		"status":   "sale",
	}
	var form = CreateProductForm{}
	BindForm(&form, extraParams, image_path)
	return form.Image
}

func GetImageFileDimension(path string) (width, height int) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0
	}
	return image.Width, image.Height
}
