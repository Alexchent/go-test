package main

import (
	"context"
	"net/http"
	"time"
)

// 传递取消信号，结束任务

func main() {
	//mock()
	//DoLater(func() {
	//	fmt.Println("延时任务开始")
	//}, 2*time.Second)

	DoAfter(2*time.Second, func() {
		println("2s时间到" +
			"延时任务开始")
	})
}

func mockCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		// 模拟子 goroutine 持续工作
		for range time.Tick(1 * time.Second) {
			select {
			case <-ctx.Done():
				println("goroutine exit")
				return
			default:
				println("goroutine running")
			}
		}
	}()

	time.Sleep(3 * time.Second)
	// 通知子 goroutine 退出
	cancel()
	// 等待子 goroutine 退出
	time.Sleep(1 * time.Second)
}

// DoAfter 传递超时信号，实现延时任务
func DoAfter(delay time.Duration, f func()) {
	// 3s 后超时
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		f()
	case <-time.After(10 * time.Second):
		println("任务超时")
	}
}

// DoLater 利用context实现一个延迟任务
func DoLater(f func(), delay time.Duration) {
	select {
	case <-time.After(delay):
		f()
	}
}

type key int

const (
	requestIDKey key = iota
)

// 传递上下文信息
func WithRequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// 从请求中提取请求ID和用户信息
		requestID := req.Header.Get("X-Request-ID")

		// 创建子 context，并添加一个请求 Id 的信息
		ctx := context.WithValue(req.Context(), requestIDKey, requestID)

		// 创建一个新的请求，设置新 ctx
		req = req.WithContext(ctx)

		// 将带有请求 ID 的上下文传递给下一个处理器
		next.ServeHTTP(rw, req)
	})
}
