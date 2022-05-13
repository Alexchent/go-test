package main

import (
	"fmt"
	"github.com/alexchen/test/Cache"
	"strings"
)

// 清理后缀是js和torrent的文件
func main() {
	key := "laravel_database_files"
	val := Cache.SMembers(key)

	for _, v := range val {
		//fmt.Println(v)
		//newv := strings.TrimRight(v, "\n")
		//fmt.Println(newv)
		//return
		if strings.HasSuffix(v, "js") || strings.HasSuffix(v, "torrent") {
			fmt.Println(v)
			Cache.SRem(key, v)
		}
	}
}
