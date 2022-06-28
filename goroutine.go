package main 

import "fmt"


//轻量级线程——协程 实现并发

func main() {
	s := []int {7,2,8,-9,4,0}
	//分治思想
    c := make(chan int)
    //println(len(s))
    //开启协程
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)

    x, y := <-c, <-c //从通道中取数据  阻塞等待通道中的数据

    fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _,v := range s {
		sum += v
	}
	//sum发送到信道c中
	c <- sum
}