package lib

import (
	"github.com/parnurzeal/gorequest"

	"net/http"
)

func (app *App) HttpRequest(method string, route string, body interface{}) (resp *http.Response, resp_body string) {
	request := gorequest.New()
	resp, resp_body, errs := request.Get(route).End()
	if errs != nil {
		PR_DUMP(errs)
		return nil, ""
	}
	return resp, resp_body
}
