package main

import "fmt"

func main() {
	//交换变量的值
	a,b := 1,2
	b,a = a,b
	fmt.Println(a,b);
}