package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	stopC := make(chan struct{})
	go func() {
		// 只有一个channel执行，如果没有则会阻塞
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("超市")
		case <-stopC:
			fmt.Println("终止")
		case <-ch:
			fmt.Println("go go go")
		}
	}()
	//ch <- 1
	stopC <- struct{}{}
	time.Sleep(2 * time.Second)
	fmt.Println("程序结束")
}

func test() {
	ch := make(chan int)
	quit := make(chan bool)

	//新开一个协程
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
				//default :
				//	fmt.Println("wait: 1s")
				//	time.Sleep(time.Second)
			}
		}
	}() //别忘了()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit //阻塞直到有输出
	fmt.Println("程序结束")
}
