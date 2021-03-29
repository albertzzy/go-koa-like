package types

// import (
// 	App "go-koa-like/lib"
// )

type NextType func() interface{}

// type MidType func(ctx *App.Context, next NextType) interface{}
type MidType func(ctx interface{}, next NextType) interface{}
