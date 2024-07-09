package main

import (
	"fmt"
	"time"
)

func main() {
	_ = StartTicker(func() {
		fmt.Println("you are good")
	}, time.Second*2)

	time.Sleep(time.Second * 3)
	fmt.Println("3 end")
}

func demo() {
	t := time.NewTimer(time.Second * 2)
	<-t.C
	fmt.Println("time 1 reached")

	t1 := time.NewTimer(time.Second)
	go func() {
		<-t1.C
		fmt.Println("time 2 reached")
	}()
	// 定时器可以在失效之前取消
	stop := t1.Stop()
	if stop {
		fmt.Println("tim2 2 stop")
	}
}

// StartTicker 定时器函数
func StartTicker(f func(), d time.Duration) chan struct{} {
	done := make(chan struct{}, 1)
	go func() {
		timer := time.NewTimer(d)
		defer timer.Stop()
		select {
		case <-timer.C:
			f()
		case <-done:
			fmt.Println("提前中止")
			return
		}
	}()
	return done
}
