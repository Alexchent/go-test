package main

import (
	"fmt"
	"github.com/alexchen/go_test/cache/redis"
	"strings"
)

// 清理后缀是js和torrent的文件
func main() {
	key := "have_save_file"
	val := redis.SMembers(key)
	for _, v := range val {
		// 按 后缀清理
		if strings.HasSuffix(v, "js") || strings.HasSuffix(v, "torrent") {
			fmt.Println(v)
			redis.SRem(key, v)
		}
	}
}
