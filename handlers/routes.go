package handlers

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
)

func GetRoutes() Routes {
	return Routes{
		Route{"GET", "/", IndexHandler},
		Route{"GET", "/products", GetAllProductsHandler},
		Route{"GET", "/products/offset={offset}&limit={limit}", GetPageProductsHandler},
		Route{"POST", "/products", CreateProductHandler},
		Route{"PUT", "/products/{id}", UpdateProductHandler},
		Route{"DELETE", "/products/{id}", DeleteProductHandler},
	}
}
