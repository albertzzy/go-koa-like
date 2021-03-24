package lib

import (
	"net/http"
)

type MidFunc func(ctx *Context) interface{}
type HandlerType func(req *http.Request, res http.ResponseWriter) interface{}

type Application struct {
	context Context
	// server  http.Server

	middleware []MidFunc
}

// create a new context
func (this *Application) CreateContext(req *http.Request, res http.ResponseWriter) *Context {
	ctx := &Context{
		res: res,
		req: req,
	}
	ctx.res = res
	ctx.req = req
	// ctx.app = this

	return ctx
}

// init http server and listen
func (this *Application) Listen(port string, args ...interface{}) {
	err := http.ListenAndServe(port, this)
	if err != nil {
		this.context.OnError(err, 500)
	}
}

// add middle ware
func (this *Application) Use(fn MidFunc) *Application {
	this.middleware = append(this.middleware, fn)
	return this
}

// http handler
func (this *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// compose fn
	fn := utils.Compose(this.middleware)
	// ctx := this.CreateContext(req, res)
	// this.handleRequest(ctx, fn)
}

func (this *Application) handleRequest(ctx *Context, fn MidFunc) interface{} {
	// res := ctx.res
	// res.statusCode = 404
	// onerror := func(err) {
	// 	this.context.OnError(err)
	// }
	// handleResponse := func() {
	// 	// respond(ctx);
	// }
	// onFinished(res, onerror);
	// return fn(ctx).then(handleResponse).catch(onerror)
}

// new  application
func New() {

}
