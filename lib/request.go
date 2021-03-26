package lib

import (
	"net/http"
	"strings"
)

type Request struct {
	req *http.Request
}

func (this *Request) Header() http.Header {
	return this.req.Header
}

func (this *Request) Get(field string) string {
	req := this.req
	if strings.ToLower(field) == "referer" || strings.ToLower(field) == "refererr" {
		return req.Header.Get("referer")
	}
	return req.Header.Get(field)
}
