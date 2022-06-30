package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// 判断变量 i 的数据类型是不是 string
	v, ok := i.(string)
	fmt.Println(v, ok)

	f, ok := i.(int)
	fmt.Println(f, ok)
}
