package utils

import (
	"errors"
	"sync"
)

/*
	import cycle
	recursion
	异步阻塞，非阻塞问题  node, go
*/

type FuncType func(ctx interface{}, next func() interface{}) interface{}

var index = -1
var wg sync.WaitGroup
var ch = make(chan interface{})

func dispatch(i int, num int, midware []FuncType, context interface{}, next FuncType) interface{} {
	if i <= index {
		return errors.New("next() called multiple times")
	}
	index = i
	fn := midware[i]
	if i == num {
		fn = next
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

func Compose(midware []FuncType) interface{} {
	return func(context interface{}, next FuncType) interface{} {
		num := len(midware)
		return dispatch(0, num, midware, context, next)
	}
}
