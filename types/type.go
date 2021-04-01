package types

import (
	. "go-koa-like/lib"
)

type NextType func() interface{}

type MidType func(ctx *Context, next NextType) interface{}
