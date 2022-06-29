package Cache

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

// 在包目录下运行 go 工具，也可以忽略包路径
// go test

// PASS
// ok      github.com/alexchen/test/Cache  0.139s
func TestSet(t *testing.T) {
	Set("test", "hello", time.Second*10)
	v := Get("test")
	if v != "hello" {
		t.Error("失败")
	}
}

func TestGet(t *testing.T) {
	val2, err := Client.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
