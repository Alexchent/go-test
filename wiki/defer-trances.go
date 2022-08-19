package main

import "fmt"

// 延迟执行，注意如果是defer func其参数是预执行的
// 嵌套的defer先进后出
// defer 延迟执行，即使发生panic

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	b()
}

// 执行结果如下
//entering: b
//in b
//entering: a
//in a
//leaving: a
//leaving: b
