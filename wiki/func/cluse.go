package main

import "fmt"

//在闭包内部修改引用的变量
//闭包对它作用域上部的变量可以进行修改，修改引用的变量会对变量进行实际修改，通过下面的例子来理解：
func main() {
	str := "hello world"
	sum := 1
	// 创建一个无参数的匿名函数
	foo := func() {
		// 改变外部变量
		str = "hello kitty"
		sum++
	}

	foo()
	fmt.Println(str) // 发现变量已发生改变

	foo()
	//被捕获到闭包中的变量让闭包本身拥有了记忆效应，闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命期一直存在，闭包本身就如同变量一样拥有了记忆效应。
	fmt.Println(sum)
}
