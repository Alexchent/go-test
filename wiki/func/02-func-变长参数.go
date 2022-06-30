package main

import "fmt"

func main() {
	Greeting("hello")
	Greeting("hello", "alex", "summer", "pin")
}

// Greeting 传递 0 或 多个 参数
func Greeting(prefix string, who ...string) {
	fmt.Println(who)
	for k, name := range who {
		fmt.Println(k, prefix, name)
	}
}
