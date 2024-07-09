package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 无缓冲 通道
	c := make(chan int)

	// 注意 无缓冲通道 接收者和生产者必须同时存在

	// 生成5个worker并发处理通道c内的数据
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		// 启动消费协程，如无数据则阻塞
		go worker.process(c)
		// worker.process(c)   如果不启动协程，会因为没有检测到生产者，而报错：all goroutines are asleep - deadlock!
	}
	// 每个50毫秒生成一个数据传入通道c内
	for {
		c <- rand.Int()
		//fmt.Println(len(c))
		time.Sleep(time.Millisecond * 50)
	}

}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 500)
	}
}
