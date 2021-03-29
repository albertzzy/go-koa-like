package lib

import (
	"net/http"
	"strings"
)

type Request struct {
	/*
		have to be uppercase:
		https://stackoverflow.com/questions/37780520/unknown-field-in-struct-literal/37780565
	*/
	Req *http.Request
}

func (this *Request) Header() http.Header {
	return this.Req.Header
}

func (this *Request) Get(field string) string {
	req := this.Req
	if strings.ToLower(field) == "referer" || strings.ToLower(field) == "refererr" {
		return req.Header.Get("referer")
	}
	return req.Header.Get(field)
}
