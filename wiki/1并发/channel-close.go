package main

import (
	"fmt"
	"sync"
)

func main() {
}

// 利用channel的关闭来进行并发周知
func controlDemo1() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 100)

	wg.Add(1)
	go func() {
		defer wg.Done()

		//for v := range ch {
		//	fmt.Println(v)
		//}
		// 等同于
		for {
			if v, ok := <-ch; ok == true {
				fmt.Println(v)
			} else {
				break
			}
		}

	}()

	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)

	wg.Wait()
	fmt.Println("finish")
}
