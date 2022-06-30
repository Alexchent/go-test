package main

import "fmt"

func main() {
	p := new(int)  // new 函数声明一个内存，返回内存指针
	fmt.Println(p) // 输出 0xc00001c090

	a := new(map[int]string)
	fmt.Println(a) // 输出 &map[]
}
