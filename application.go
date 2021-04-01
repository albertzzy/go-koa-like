package app

import (
	"encoding/json"
	. "go-koa-like/lib"
	. "go-koa-like/types"
	utils "go-koa-like/utils"
	"log"
	"net/http"
	"strconv"
)

// type NextType func() interface{}
// type MidType func(ctx *Context, next NextType) interface{}
type HandlerType func(req *http.Request, res http.ResponseWriter) interface{}

type Application struct {
	context *Context

	middleware []MidType
}

// create a new context
func (this *Application) CreateContext(req *http.Request, res http.ResponseWriter) *Context {
	request := &Request{
		Req: req,
	}

	response := &Response{
		Res: res,
	}

	ctx := &Context{
		Res:      res,
		Req:      req,
		Request:  request,
		Response: response,
	}
	// ctx.app = this

	return ctx
}

// init http server and listen
func (this *Application) Listen(port string) {
	err := http.ListenAndServe(port, this)
	if err != nil {
		this.context.OnError(err, 500)
	}
}

// add middle ware
func (this *Application) Use(fn MidType) *Application {
	this.middleware = append(this.middleware, fn)
	return this
}

// implement the Handler interface
// !!!servehttp执行了两次，一次是 / 一次是 /favicon.ico
func (this *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// compose fn
	fn := utils.Compose(this.middleware)
	ctx := this.CreateContext(req, res)
	this.handleRequest(ctx, fn)
}

func (this *Application) handleRequest(ctx *Context, fn MidType) {
	// res := ctx.res
	// res.statusCode = 404
	// onerror := func(err error) {
	// 	this.context.OnError(err, 404)
	// }
	// onFinished(res, onerror);

	fn(ctx, func() interface{} {
		return Respond(ctx)
	})

}

// new  application
func New() *Application {
	return &Application{}
}

// handle reponse
func Respond(ctx *Context) interface{} {
	res := ctx.Res
	body := ctx.Body

	statusCode, err := strconv.Atoi(ctx.Response.Status())
	log.Println("====handlerequest=====", statusCode, err)
	if err != nil {
		// todo - onerror
		return ctx
	}

	res.WriteHeader(statusCode)

	// body: string | json
	resBody, ok := body.(string)

	log.Println(resBody, ok)

	if ok {
		res.Write([]byte(resBody))
	} else {
		resp, err2 := json.Marshal(body)
		if err2 != nil {
			// todo - onerror
			return ctx
		}
		res.Write(resp)
	}

	return ctx
}
