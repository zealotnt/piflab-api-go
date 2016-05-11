package factory

import (
	"bytes"
	"github.com/icrowley/fake"
	. "github.com/mitchellh/mapstructure"
	"github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"
	"math/rand"
	"os"
	"time"
)

func NewProduct(params ...map[string]interface{}) (*Product, error) {
	rand.Seed(time.Now().UTC().UnixNano())

	fh, err := os.Open("db/seeds/factory/golang.png")
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	dataBytes := bytes.Buffer{}

	dataBytes.ReadFrom(fh)

	product := &Product{
		Name:      fake.ProductName(),
		Price:     rand.Intn(100000),
		Provider:  fake.Company(),
		Rating:    rand.Float32() * float32(rand.Intn(5)),
		Status:    STATUS_OPTIONS[rand.Intn(len(STATUS_OPTIONS))],
		Detail:    fake.ParagraphsN(1),
		ImageData: dataBytes.Bytes(),
		Image:     "golang.png",
	}

	if params != nil {
		err := Decode(params[0], product)
		return product, err
	}

	return product, nil
}

func CreateProduct(DB *lib.DB, params ...map[string]interface{}) (*Product, error) {
	product, err := NewProduct(params...)

	if err != nil {
		return product, err
	}

	return product, (ProductRepository{DB}).SaveProduct(product)
}
