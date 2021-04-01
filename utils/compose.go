package utils

import (
	"errors"
	. "go-koa-like/lib"

	. "go-koa-like/types"
)

/*
	import cycle
	recursion
	异步阻塞，非阻塞问题  node, go
	中间件同步执行
*/

var index = -1

// var wg sync.WaitGroup
// var ch = make(chan interface{}, 100)

func Compose(midware []MidType) MidType {
	return func(context *Context, next NextType) interface{} {
		index = -1
		num := len(midware)
		res := dispatch(0, num, midware, context, next)
		// wg.Wait()
		return res
	}
}

func dispatch(i int, num int, midware []MidType, context *Context, next NextType) interface{} {
	if i <= index {
		return errors.New("next() called multiple times")
	}

	index = i

	var fn MidType
	if i < num {
		fn = midware[i]
	} else if i == num {
		// last Response(ctx)
		return next()
	} else {
		return nil
	}

	/*
		//wrong version:
		wg.Add(1)
		go fn(context, func() interface{} {
			log.Println("==dispatch index====", i)

			res := dispatch(i+1, num, midware, context, next)
			log.Println("==dispatch return====", res)
			wg.Done()
			ch <- res
			return res
		})
		return <-ch
	*/

	// syncronuous
	return fn(context, func() interface{} {
		return dispatch(i+1, num, midware, context, next)
	})

}
