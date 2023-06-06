package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(1 * time.Second)
	defer timer.Stop()

	// 定时执行
	for range timer.C {
		fmt.Println("Timer fired", time.Now().Unix())
	}
}
