package factory

import (
	"github.com/icrowley/fake"
	. "github.com/mitchellh/mapstructure"
	"github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"
	"math/rand"
	"time"
)

var status = []string{
	"sale",
	"out of stock",
	"available",
}

func NewProduct(params ...map[string]interface{}) *Product {
	rand.Seed(time.Now().UTC().UnixNano())

	id := fake.CharactersN(14)
	product := Product{
		Name:     fake.ProductName(),
		Price:    rand.Intn(100000),
		Provider: fake.Company(),
		Rating:   rand.Float32() * float32(rand.Intn(5)),
		Status:   status[rand.Intn(len(status))],
		Image:    id + ".jpg",
		Detail:   id,
	}

	if params != nil {
		if err := Decode(params[0], &product); err != nil {
			panic(err)
		}
	}

	return &product
}

func CreateProduct(app *lib.App, params ...map[string]interface{}) *Product {
	product := NewProduct(params...)

	err := ProductRepository{app.DB}.CreateProduct(product)

	if err != nil {
		panic(err)
	}

	return product
}

func CountProduct(app *lib.App) int {
	products, err := ProductRepository{app.DB}.GetAll()

	if err != nil {
		panic(err)
	}

	return len(*products)
}
