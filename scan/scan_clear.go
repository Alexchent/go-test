package main

import (
	"fmt"
	"github.com/alexchen/go_test/Cache/redis"
	"log"
	path "path/filepath"
	"strings"
)

// 清理后缀是js和torrent的文件
func main() {
	key := "have_save_file"
	val := redis.SMembers(key)
	for _, v := range val {
		//fmt.Println(v)
		// 清理末尾有换行符的值
		fileExt := path.Ext(v)
		//fmt.Println(fileExt)
		log.Print(fileExt)

		newv := strings.TrimRight(v, "\n")
		redis.SRem(key, v)
		redis.SAdd(key, newv)

		//fmt.Println(newv)
		//return
		// 按 后缀清理

		if strings.HasSuffix(v, "js") || strings.HasSuffix(v, "torrent") {
			fmt.Println(v)
			redis.SRem(key, v)
		}
	}
}
