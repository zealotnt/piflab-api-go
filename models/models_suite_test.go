package models_test

import (
	"github.com/o0khoiclub0o/piflab-store-api-go/db/seeds/factory"
	. "github.com/o0khoiclub0o/piflab-store-api-go/handlers"
	"github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
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

	By("Automatically create some non-image products")
	sJson := `{"no-image": "yes"}`
	extraParams := make(map[string]interface{})
	factory.Json2Map(sJson, extraParams)
	for i := 0; i < 10; i++ {
		_, err := factory.CreateProduct(app.DB, extraParams)
		Expect(err).To(BeNil())
	}
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
