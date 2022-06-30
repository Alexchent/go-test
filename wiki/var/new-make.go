package main

import (
	"fmt"
	"time"
)

func main() {
	// new 声明一个变量，并分配内存，返回值是这个变量的指针
	p := new(int)  // new 函数声明一个内存，返回内存指针
	fmt.Println(p) // 输出 0xc00001c090

	a := new(map[int]string)
	fmt.Println("new map 返回值：", a) // 输出 &map[]

	// make 初始化map、slice、chan类型变量，并分配内存。返回对象的引用
	sli := make([]int, 3)
	fmt.Println("make array 返回值", sli)

	map1 := make(map[int]string, 10)
	fmt.Println("make map 返回值：", map1)

	// make channel
	ch := make(chan string, 10)
	go func() {
		fmt.Println(<-ch)
	}()

	ch <- "hello"
	time.Sleep(time.Second)
}
