package main

import "fmt"

func main() {
	var name string

	fmt.Printf("请输入您的名字: ")

	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}

	fmt.Printf("hello %s", name)

}
