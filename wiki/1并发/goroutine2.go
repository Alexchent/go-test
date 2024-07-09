package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 2; i++ {
		i := i
		go func() {
			defer fmt.Println("worker", i, "exit")
			for v := range c {
				fmt.Printf("worker %d got %d\n", i, v)
				//time.Sleep(time.Millisecond * 500)
			}
		}()
	}
	// 每个50毫秒生成一个数据传入通道c内
	c <- 1
	c <- 2
	c <- 3
	time.Sleep(time.Second * 1)
	close(c)
	time.Sleep(time.Second * 1)
}
