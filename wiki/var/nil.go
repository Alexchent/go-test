package main

import "fmt"

func main() {
	// bool类型，数值类型，字符串类型 的默认值都是对应的零值 false 0 ""，不是nil
	// 指针，channel，函数，interface，slice，map的零值都是nil

	var a interface{} //var 声明 interface，值为nil
	var b interface{}
	fmt.Println(a == b) //true

	var c []int //var 声明 slice，值为nil
	var d []int
	//fmt.Println(c == d) //invalid operation: c == d (slice can only be compared to nil)
	fmt.Println(c == nil) //true
	fmt.Println(d == nil) //true

	for range []int(nil) { //循环次数将是0
		fmt.Println("Hello")
	}

	for range map[string]string(nil) { //循环次数将是0
		fmt.Println("world")
	}

	for i := range (*[5]int)(nil) {
		fmt.Println(i) // 0 1 2 3 4
	}

	for range chan bool(nil) { // block here
		fmt.Println("Bye") //fatal error: all goroutines are asleep - deadlock!
	}
}
