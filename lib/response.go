package lib

type Response struct{}

func (this *Response) header() interface{} {
	return this.res.getHeaders()
}

func (this *Request) status() string {
	return this.res.statusCode
}

func (this *Response) set(args ...interface{}) {
	if len(args) == 2 {
		this.res.setHeader(args[0], args[1])
	} else {
		for key := range args {
			this.set(key, args[key])
		}
	}
}
