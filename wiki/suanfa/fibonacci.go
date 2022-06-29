package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer fmt.Println(time.Now().Sub(start))

	for i := 1; i <= 41; i++ {
		result := fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

// 打印斐波那契数列
func fibonacci(n int) (res int) {
	if n < 2 {
		res = n
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
