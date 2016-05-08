package lib

import (
	"net/url"
	"strconv"
)

type Context struct {
	Params    map[string]string
	GetParams url.Values
}

func (c Context) ID() uint {
	id, err := strconv.ParseUint(c.Params["id"], 10, 32)

	if err != nil {
		return 0
	}

	return uint(id)
}
