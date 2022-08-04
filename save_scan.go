package main

import (
	"fmt"
	"github.com/alexchen/go_test/Cache/redis"
	file "github.com/alexchen/go_test/File"
	scan "github.com/alexchen/go_test/ScanService"
	"strings"
	"time"
)

// 备份缓存中的数据到本地
func main() {
	start := time.Now()
	defer fmt.Println(time.Since(start))

	var data []string
	filename := fmt.Sprintf("./log/"+scan.SavePath, time.Now().Format("060102"))
	data = redis.SMembers("have_save_file")
	for _, v := range data {
		//scan.AppendContent(strings.TrimRight(v, "\n") + "\n")
		file.AppendContent(filename, strings.Trim(v, "\n")+"\n")
	}

	data = redis.SMembers("laravel_database_files")
	for _, v := range data {
		file.AppendContent(filename, v+"\n")
	}
}
