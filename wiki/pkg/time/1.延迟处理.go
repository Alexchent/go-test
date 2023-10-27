package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	DoAfter2(time.Second*2, func() {
		fmt.Println("延时任务开始")
	})
}

func DoAfter(delay time.Duration, f func()) {
	select {
	case <-time.After(delay):
		f()
	}
}

func DoAfter2(delay time.Duration, f func()) {
	ctx, cancel := context.WithTimeout(context.Background(), delay)
	defer cancel()

	select {
	case <-ctx.Done():
		f()
	}
}
