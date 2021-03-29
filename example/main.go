package main

import (
	"fmt"
	// . "go-koa-like/lib"
	// . "go-koa-like/types"
)

var ch = make(chan int, 1)

func dispatchs(i int) {
	<-ch
	if i > 5 {
		return
	}
	fmt.Println(i)
	ch <- 1
	go dispatchs(i + 1)
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
	} */

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
	// dispatchs(0)

	/* app := New()
	app.Use(func(ctx *Context, next NextType) interface{} {
		ctx.body = "Hello world"
		return nil
	})
	app.Listen(":9001") */

}
