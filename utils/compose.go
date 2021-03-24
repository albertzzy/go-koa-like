package lib

import (
	"errors"
)

/*
	import cycle
	recursion
	异步阻塞，非阻塞问题  node, go
*/

type FuncType func(interface{}, func()) interface{}

func Compose(midware []FuncType) interface{} {

	return func(context interface{}, next FuncType) interface{} {
		index := -1
		ch := make(chan int, 1)
		dispatch := func(i int) interface{} {
			if i <= index {
				return errors.New("next() called multiple times")
			}
			index = i
			fn := midware[i]
			if i == len(midware) {
				fn = next
			}

			ch <- 1
			go fn(context, func() {
				<-ch
				return dispatch(i + 1)
			})
		}

		return dispatch(0)

	}
}
