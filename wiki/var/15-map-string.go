package main

import (
	"fmt"
)

func main() {
	// go语言中字符串是不可变的，也就是说 str[index] 这样的表达式是不可以被放在等号左侧的
	// 否则会得要错误：cannot assign to str[i]
	s := "hello"
	// 字符串转数组
	c := []byte(s)
	c[0] = 'c'
	// 数组转字符串
	s2 := string(c)
	fmt.Println(s2)
}
