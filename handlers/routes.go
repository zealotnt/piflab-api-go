package handlers

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
)

func GetRoutes() Routes {
	return Routes{
		Route{"GET", "/", IndexHandler},
		Route{"GET", "/products", GetProductsHandler},
		Route{"GET", "/products/{id}", GetProductsDetailHandler},
		Route{"POST", "/products", CreateProductHandler},
		Route{"PUT", "/products/{id}", UpdateProductHandler},
		Route{"DELETE", "/products/{id}", DeleteProductHandler},

		Route{"GET", "/cart", GetCartHandler},
		Route{"PUT", "/cart/items", UpdateCartHandler},
		Route{"PUT", "/cart/items/{id}", UpdateCartItemHandler},
		Route{"DELETE", "/cart/items/{id}", DeleteCartItemHandler},

		Route{"POST", "/cart/checkout", CheckoutCartHandler},
		Route{"GET", "/orders", GetCheckoutHandler},
		Route{"GET", "/orders/{id}", GetCheckoutDetailHandler},

		Route{"OPTIONS", "", OptionHandler},
	}
}
