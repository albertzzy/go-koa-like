package lib

import (
	utils "go-koa-like/utils"
	"net/http"
)

// type MidFunc func(ctx *Context, next func() interface{}) interface{}
type MidFunc func(ctx interface{}, next func() interface{}) interface{}
type HandlerType func(req *http.Request, res http.ResponseWriter) interface{}

type Application struct {
	context Context

	middleware []utils.FuncType
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
func (this *Application) Use(fn utils.FuncType) *Application {
	this.middleware = append(this.middleware, fn)
	return this
}

// implement the Handler interface
func (this *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// compose fn
	fn := utils.Compose(this.middleware)
	ctx := this.CreateContext(req, res)
	this.handleRequest(ctx, fn)
}

func (this *Application) handleRequest(ctx *Context, fn interface{}) interface{} {
	// res := ctx.res
	// res.statusCode = 404
	onerror := func(err error) {
		this.context.OnError(err, 404)
	}
	// onFinished(res, onerror);
	fn(ctx, func() {})

	utils.Respond(ctx)
}

// new  application
func New() *Application {

	return &Application{}
}
