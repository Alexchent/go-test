package main

import "fmt"

func main() {
	one := 0
	one, two := 1, 2 // 你不能在一个单独的声明中重复声明一个变量，但在多变量声明中这是允许的，其中至少要有一个新的声明变量。

	fmt.Println(one, two)
}
