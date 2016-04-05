package main

import (
	"fmt"
	"github.com/o0khoiclub0o/piflab-store-api-go/db/seeds/factory"
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
)

func main() {
	app := NewApp()
	defer app.Close()

	if factory.CountProduct(app) == 0 {
		for i := 0; i < 10; i++ {
			factory.CreateProduct(app)
		}
	}

	fmt.Println("Seed successfully")
}
