package main

import (
	"fmt"
	"github.com/alexchen/go_test/Cache/redis"
	"time"
)

func main() {
	// 格式化时间
	fmt.Println(time.Now().Format("20060102"))

	// 测试redis
	redis.Set("name", "baobo", time.Second*10)

	fmt.Println(redis.Get("name"))
}
