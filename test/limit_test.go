package main

import (
	"fmt"
	"github.com/alexchen/go_test/util"
	"sync"
	"testing"
	"time"
)

func TestCounterLimit(t *testing.T) {
	var wg sync.WaitGroup
	ls := util.NewCounterLimit(1*time.Second, 5)

	for i := 0; i < 7; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if ls.Allow() {
				println("ok")
			} else {
				println("no")
			}
		}()
	}

	time.Sleep(2 * time.Second)

	fmt.Println("after 2s")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if ls.Allow() {
				println("ok")
			} else {
				println("no")
			}
		}()
	}

	wg.Wait()
}
