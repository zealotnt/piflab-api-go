package models_test

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/handlers"
	"github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}

var app *lib.App

var _ = BeforeSuite(func() {
	app = lib.NewApp()
	app.AddRoutes(GetRoutes())
})

var _ = AfterSuite(func() {
	app.Close()
})

func Request(method string, route string, body string) *httptest.ResponseRecorder {
	return app.Request(method, route, body)
}

func getProducts(body []byte) (*[]Product, error) {
	products := &[]Product{}
	if err := json.Unmarshal(body, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func getFirstAvailableId(response *httptest.ResponseRecorder) uint {
	body, _ := ioutil.ReadAll(response.Body)
	products, _ := getProducts(body)

	for idx := range *products {
		return (*products)[idx].Id
	}

	return 0
}

func getFirstAvailableUrl() string {
	response := Request("GET", "/products", "")
	return fmt.Sprintf("/products/%d", getFirstAvailableId(response))
}

func RequestPost(method string, route string) *httptest.ResponseRecorder {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, _ := bodyWriter.CreateFormFile("image", "golang.png")
	fh, _ := os.Open(os.Getenv("FULL_IMPORT_PATH") + "/db/seeds/factory/golang.png")
	io.Copy(fileWriter, fh)
	bodyWriter.WriteField("name", "xbox")
	bodyWriter.WriteField("price", "70000")
	bodyWriter.WriteField("provider", "Microsoft")
	bodyWriter.WriteField("rating", "3.5")
	bodyWriter.WriteField("status", "sale")
	bodyWriter.Close()

	request, _ := http.NewRequest(method, route, bodyBuf)

	request.RemoteAddr = "127.0.0.1:8080"
	contentType := bodyWriter.FormDataContentType()
	request.Header.Set("Content-Type", contentType)

	response := httptest.NewRecorder()
	app.ServeHTTP(response, request)

	return response
}
