package lib

import (
	"net/http"
)

type Context struct {
	http.CookieJar

	res http.ResponseWriter
	req *http.Request
}

func (this *Context) OnError(err error, code int) {
	http.Error(this.res, err.Error(), code)
}

func (this *Context) Throw() {

}
