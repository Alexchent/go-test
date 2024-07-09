package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "process job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 启动3个worker
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(&wg, i, jobs, results)
	}

	// stage1 send发
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		for r := range results {
			fmt.Println(r)
		}
	}()

	// worker是results channel的sender，因此，等所有worker完成后关闭channel
	wg.Wait()
	close(results)
}
