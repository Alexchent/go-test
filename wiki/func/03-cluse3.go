package main

import "fmt"

func main() {
	jia := Adder()

	fmt.Println(jia(1))
	fmt.Println(jia(10))
	//fmt.Println(jia(100))
}

// Adder 匿名函数作为返回值的函数，构成闭包。 闭包的特点是在生命周期内，内存空间不会释放。
func Adder() func(int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}
