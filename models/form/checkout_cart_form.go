package models

// import (
// 	"github.com/mholt/binding"
// 	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
// 	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
// 	. "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"

// 	"errors"
// 	"net/http"
// )

// type CheckoutCartForm struct {
// 	Product_Id  *uint   `json:"product_id"`
// 	Quantity    *int    `json:"quantity"`
// 	AccessToken *string `json:"access_token"`
// }

// func (form *CheckoutCartForm) FieldMap(req *http.Request) binding.FieldMap {
// 	return binding.FieldMap{
// 		&form.Product_Id: binding.Field{
// 			Form: "product_id",
// 		},
// 		&form.Quantity: binding.Field{
// 			Form: "quantity",
// 		},
// 		&form.AccessToken: binding.Field{
// 			Form: "access_token",
// 		},
// 	}
// }

// func (form *CheckoutCartForm) Validate(method string) error {

// }
