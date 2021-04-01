package lib

import (
	"net/http"
	"net/url"
	"time"
)

type Context struct {
	Res      http.ResponseWriter
	Req      *http.Request
	Request  *Request
	Response *Response
	Body     interface{}
}

func (this *Context) OnError(err error, code int) {
	http.Error(this.Res, err.Error(), code)
}

/*
cookie: refer to iris
*/
var (
	CookieExpireDelete = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
)

func (this *Context) SetCookie(c *http.Cookie) {
	http.SetCookie(this.Res, c)
}

func (this *Context) DelCookie(name string) {
	c := &http.Cookie{}
	c.Name = name
	c.Value = ""
	c.Path = "/"
	c.HttpOnly = true
	c.Expires = CookieExpireDelete
	c.MaxAge = -1

	http.SetCookie(this.Res, c)
}

func (this *Context) GetCookie(name string) string {
	c, err := this.Req.Cookie(name)
	if err != nil {
		return ""
	}
	value, _ := url.QueryUnescape(c.Value)
	return value
}

func (this *Context) Throw() {

}
