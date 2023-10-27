package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(1 * time.Second)
	defer timer.Stop()

	quit := make(chan struct{})
	// 定时执行
	go func() {
		defer fmt.Println("Timer1 goroutine exit")
		for {
			select {
			case <-timer.C:
				fmt.Println("Timer1 fired", time.Now().Unix())
			case <-quit:
				fmt.Println("Timer1 quit")
				return
			}
		}
	}()
	go func() {
		defer fmt.Println("Timer2 goroutine exit")
		for {
			select {
			case <-timer.C:
				fmt.Println("Timer2 fired", time.Now().Unix())
			case <-quit:
				fmt.Println("Timer2 quit")
				return
			default:
				fmt.Println("Timer2 running", time.Now().Unix())
			}
		}
	}()
	fmt.Println("Timer start")
	time.Sleep(3 * time.Second)

	fmt.Println("Timer close")
	//quit <- struct{}{} // 关闭信号可以用于关闭单个定时器
	close(quit) // 关闭信号可以用于关闭多个定时器
	for {
		//fmt.Println("main goroutine running")
	}
}
