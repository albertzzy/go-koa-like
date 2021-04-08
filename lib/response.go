package lib

import (
	"net/http"
	"strconv"
)

type Response struct {
	Res    http.ResponseWriter
	values map[string]interface{}
}

func (this *Response) Status() string {
	code, ok := this.values["statusCode"].(int)
	if !ok {
		//
		return "500"
	}
	return strconv.Itoa(code)
}

func (this *Response) Set(args ...interface{}) {
	if len(args) == 2 {
		arg0, ok0 := args[0].(string)
		arg1, ok1 := args[1].(string)
		if !ok0 || !ok1 {
			// todo - ctx.onError
			return
		}
		this.values[arg0] = arg1
		this.Res.Header().Set(arg0, arg1)
	} else {
		for key := range args {
			this.Set(key, args[key])
		}
	}
}

func (this *Response) StatusCode(code int) {
	this.values["statusCode"] = code
	this.Res.WriteHeader(code)
}

func (this *Response) NotFound() {
	this.StatusCode(http.StatusNotFound)
}

func (this *Response) Get(field string) string {
	return this.Res.Header().Get(field)
}
