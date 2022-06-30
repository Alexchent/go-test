package main

import "fmt"

func main() {
	add_func := add(1, 2)
	fmt.Println(add_func(1, 1)) // 1,3,2
	fmt.Println(add_func(0, 0)) // 2,3,0
	fmt.Println(add_func(2, 2)) // 3,3,4
}

// 闭包使用方法
func add(x1, x2 int) func(x3 int, x4 int) (int, int, int) {
	i := 0
	return func(x3 int, x4 int) (int, int, int) {
		i++
		return i, x1 + x2, x3 + x4
	}
}
