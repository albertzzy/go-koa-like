package lib

import "net/http"

type Response struct {
	res http.ResponseWriter
}

func (this *Response) Status() string {
	return this.res.Header().Get("statusCode")
}

func (this *Response) Set(args ...interface{}) {
	if len(args) == 2 {
		arg0, ok0 := args[0].(string)
		arg1, ok1 := args[1].(string)
		if !ok0 || !ok1 {
			// todo - ctx.onError
			return
		}
		this.res.Header().Set(arg0, arg1)
	} else {
		for key := range args {
			this.Set(key, args[key])
		}
	}
}

func (this *Response) Get(field string) string {
	return this.res.Header().Get(field)
}
