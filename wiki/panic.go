package main

import "fmt"

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("error 中断进程")

	fmt.Println("出现异常后，我就不会出现了")
}

// painc() 恐慌流程，终止程序，抛出异常，但不会影响前面 defer 的执行
