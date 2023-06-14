package main

import (
	"github.com/alexchen/go_test/wiki/demo"
	"time"
)

func main() {
	go demo.Slog.Send()

	for {
		demo.Slog.PreSend([]byte("hello"))
		time.Sleep(time.Second)
	}
}
