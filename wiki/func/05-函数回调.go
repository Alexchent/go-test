package main

import "fmt"

func main() {
	s := callback(1, Add)
	fmt.Println(s)

	p := callback(12, mul)
	fmt.Println(p)

	s = callback2(2, 3, Add)
	fmt.Println(s)

	p = callback2(2, 3, mul)
	fmt.Println(p)
}

// 回调函数，最后一个参数是函数类型，这个函数可以直接使用前面传入的参数
func callback(y int, f func(int, int) int) int {
	return f(y, 10)
}

// Jisuan 声明函数类型变量
type Jisuan func(a, b int) int

func callback2(x, y int, jisuan Jisuan) int {
	return jisuan(x, y)
}

func Add(x, y int) int {
	fmt.Printf("%d + %d = %d", x, y, x+y)
	res := x + y
	return res
}

func mul(a, b int) int {
	return a - b
}
