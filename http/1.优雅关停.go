package main

import (
	"context"
	"fmt"
	"github.com/alexchen/go_test/http/middleware"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()

	go func() {
		for {
			select {
			case <-time.After(time.Second * 5):
				log.Printf("datas-len: ")
			}
		}
	}()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("index"))
	})
	mux.HandleFunc("/hello", middleware.Chain(sayBye, middleware.Logging(), middleware.Recover()))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	// 创建系统信号接收器
	done := make(chan os.Signal)
	// 监听指定信号 ctrl+c kill
	// 注意kill -9 不行哦
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		// 优雅关停
		<-done
		fmt.Println("server will shutdown")

		err := server.Shutdown(context.Background())
		if err != nil {
			log.Fatal("server shutdown:", err)
		}
	}()

	fmt.Println("server run in 8080")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			fmt.Println("server closed under request")
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println("server exit")
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
