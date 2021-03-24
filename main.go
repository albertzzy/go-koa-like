package main

import (
	"fmt"
)

func main() {
	/* var wg sync.WaitGroup

	// 开N个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}

	// 等待N个后台线程完成
	wg.Wait() */

	// done := make(chan int, 10) // 带 10 个缓存

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

	ch := make(chan int, 1)
	// for i := 0; i < 5; i++ {
	// 	// ch <- i
	// 	go func(i int) {
	// 		fmt.Println(i)
	// 	}(i)
	// }

	func dispatchs(i int) {
	// dispatchs := func(i int) {
		if i > 5 {
			return
		}
		fmt.Println(i)
		go dispatchs(i + 1)
	}

	dispatchs(0)

}
