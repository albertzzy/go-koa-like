package main

import (
	"fmt"
	App "go-koa-like"
	. "go-koa-like/lib"
	. "go-koa-like/types"
	"sync"
)

var ch = make(chan interface{}, 5)
var wg sync.WaitGroup

func dispatchs(i int) {
	if i > 5 {
		return
	}
	fmt.Println(i)
	wg.Add(1)
	go func(i int) {
		dispatchs(i + 1)
		wg.Done()
		// ch <- res
	}(i)
	// return <-ch
}

func main() {
	/* var wg sync.WaitGroup

	// 开N个后台打印线程
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
		wg.Wait()
	}
	*/
	// 等待N个后台线程完成

	// done := make(chan int, 5) // 带 10 个缓存

	// // 开N个后台打印线程
	// for i := 0; i < cap(done); i++ {
	// 	go func(i int) {
	// 		// fmt.Println(i)
	// 		done <- i
	// 	}(i)
	// }

	// // 等待N个后台线程完成
	// for i := 0; i < cap(done); i++ {
	// 	a := <-done
	// 	fmt.Println(a)
	// }

	// for i := 0; i < 5; i++ {
	// 	// ch <- i
	// 	go func(i int) {
	// 		fmt.Println(i)
	// 	}(i)
	// }
	// ch <- 1

	app := App.New()
	app.Use(func(ctx *Context, next NextType) interface{} {
		ctx.Body = "Hello world"
		return next()
	})

	app.Use(func(ctx *Context, next NextType) interface{} {
		dispatchs(0)
		wg.Wait()
		return next()
	})
	app.Listen(":9001")

}
