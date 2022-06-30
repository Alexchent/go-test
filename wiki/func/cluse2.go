package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 匿名函数作为变量
	sum := 0
	Add := func(x, y int) int {
		s := x + y
		sum += s
		return sum
	}

	fmt.Println(Add(1, 2)) // 5
	fmt.Println(Add(2, 3)) // 8 记忆效应 sum变量一直保存在内存中

	// 打印匿名函数的类型
	fmt.Println(reflect.TypeOf(Add))
	fmt.Println(reflect.TypeOf(Add(1, 2)))
}
