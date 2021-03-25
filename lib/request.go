package lib

import (
	"strings"
)

type Request struct{}

func (this *Request) header() interface{} {
	return this.req.headers
}

func (this *Request) get(field string) string {
	req = this.req
	if strings.ToLower(field) == "referer" || strings.ToLower(field) == "refererr" {
		return req.headers.referer
	}
	return req.headers[field]
}
