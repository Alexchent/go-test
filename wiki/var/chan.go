package main

import "fmt"

func main() {
	//初始化一个chan c
	c := make(chan int)
	defer close(c)
	//执行一个将7发送到chan c中的方法
	go func() { c <- 3 + 4 }()
	//阻塞直到从chan c中接收数据，赋值给变量i
	i := <-c
	fmt.Println(i)
}
