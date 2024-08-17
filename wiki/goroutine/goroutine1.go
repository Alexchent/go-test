package main

import "fmt"

func main() {
	// 生产者在前
	nums := make(chan int)
	go func() {
		// 生产者必须关闭channel
		defer close(nums)
		for i := 1; i <= 100; i++ {
			nums <- i
		}
	}()

	res := 0
	// 消费者不断从channel中读取数据，当channel中无数据会阻塞直到channel关闭
	for num := range nums {
		res += num
	}

	fmt.Println(res)
}
