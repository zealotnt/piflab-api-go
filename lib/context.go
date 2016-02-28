package lib

import (
	"net/url"
)

type Context struct {
	Params    map[string]string
	GetParams url.Values
}
