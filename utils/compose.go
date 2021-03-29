package utils

import (
	"errors"
	. "go-koa-like/lib"

	// . "go-koa-like/types"
	"sync"
)

/*
	import cycle
	recursion
	异步阻塞，非阻塞问题  node, go
*/

type NextType func() interface{}
type MidType func(ctx *Context, next NextType) interface{}

var index = -1
var wg sync.WaitGroup
var ch = make(chan interface{})

func dispatch(i int, num int, midware []MidType, context *Context, next NextType) interface{} {
	if i <= index {
		return errors.New("next() called multiple times")
	}
	index = i
	fn := midware[i]
	if i == num {
		//fn is MidType, next is NextType， so wrap it
		fn = func(ctx *Context, nextFunc NextType) interface{} {
			return next()
		}
	}
	wg.Add(1)
	go fn(context, func() interface{} {
		res := dispatch(i+1, num, midware, context, next)
		wg.Done()
		ch <- res
		return res
	})
	wg.Wait()
	return <-ch
}

func Compose(midware []MidType) MidType {
	return func(context *Context, next NextType) interface{} {
		num := len(midware)
		return dispatch(0, num, midware, context, next)
	}
}
