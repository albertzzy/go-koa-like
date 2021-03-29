package lib

import (
	"net/http"
)

type Context struct {
	http.CookieJar

	Res      http.ResponseWriter
	Req      *http.Request
	Request  *Request
	Response *Response
	Body     interface{}
}

func (this *Context) OnError(err error, code int) {
	http.Error(this.Res, err.Error(), code)
}

func (this *Context) Throw() {

}
