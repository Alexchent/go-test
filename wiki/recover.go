package main

import "fmt"

func main() {
	// 从 panic 中 恢复
	defer func() {
		fmt.Println("recovered:", recover())
	}()
	panic("not good")
}
