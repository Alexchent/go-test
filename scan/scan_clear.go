package main

import (
	"fmt"
	"github.com/alexchen/go_test/Cache/redis"
	"strings"
)

// 清理后缀是js和torrent的文件
func main() {
	key := "laravel_database_files"
	val := redis.SMembers(key)

	for _, v := range val {
		//fmt.Println(v)
		//newv := strings.TrimRight(v, "\n")
		//fmt.Println(newv)
		//return
		if strings.HasSuffix(v, "js") || strings.HasSuffix(v, "torrent") {
			fmt.Println(v)
			redis.SRem(key, v)
		}
	}
}
