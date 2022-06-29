package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	//迭代 数组array、切片slice、通道channel或集合map的元素
	for index, value := range nums {
		sum += value
		fmt.Println(index, value)
	}
	fmt.Println("sum:", sum)

	ch := make(chan int)
	go func() {
		// 迭代 chan
		for c := range ch {
			fmt.Println(c)
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- i
	}

	time.Sleep(time.Second)

	map1 := make(map[int]string)
	map1 = map[int]string{1: "php", 2: "golang", 3: "python"}
	fmt.Println(map1)
	// 迭代 map
	for i, v := range map1 {
		fmt.Println(i, v)
	}
}
