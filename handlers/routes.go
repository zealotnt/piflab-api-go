package handlers

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
)

func GetRoutes() Routes {
	return Routes{
		Route{"GET", "/", IndexHandler},
		Route{"GET", "/products", GetProductsHandler},
		Route{"POST", "/products", CreateProductHandler},
		Route{"PUT", "/products/{id}", UpdateProductHandler},
		Route{"DELETE", "/products/{id}", DeleteProductHandler},
		Route{"OPTIONS", "", OptionHandler},
		Route{"GET", "/cart", GetCartHandler},
		Route{"PUT", "/cart/items", UpdateCartHandler},
	}
}
