package models

import (
	"github.com/mholt/binding"
	"net/http"
)

type GetProductForm struct {
	Offset uint
	Limit  uint
	Fields string
	Search string
}

func (form *GetProductForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&form.Offset: binding.Field{
			Form: "offset",
		},
		&form.Limit: binding.Field{
			Form: "limit",
		},
		&form.Fields: binding.Field{
			Form: "fields",
		},
		&form.Search: binding.Field{
			Form: "q",
		},
	}
}
