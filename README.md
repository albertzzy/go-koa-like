# a go version of koa

## Usage

```go
    var app = App.New()

    app.use(func(ctx *Context, next interface{}) interface{}{
        ctx.Body = "hello world"
    })

    app.Listen(":9000")
```