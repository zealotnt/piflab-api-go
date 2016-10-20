package lib

import (
	"github.com/parnurzeal/gorequest"

	"net/http"
)

func (app *App) HttpRequest(method string, route string, body interface{}) (*http.Response, string) {
	var resp *http.Response
	var resp_body string
	var errs []error

	request := gorequest.New()
	if method == "GET" {
		resp, resp_body, errs = request.Get(route).End()
	} else if method == "POST" {
		resp, resp_body, errs = request.Post(route).SendStruct(body).End()
	} else if method == "PUT" {
		resp, resp_body, errs = request.Put(route).SendStruct(body).End()
	} else if method == "DELETE" {
		resp, resp_body, errs = request.Delete(route).End()
	}
	if errs != nil {
		PR_DUMP(errs)
		return nil, ""
	}

	return resp, resp_body
}
