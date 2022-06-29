package main

import (
	"fmt"
	"github.com/alexchen/test/Cache"
	file "github.com/alexchen/test/File"
	scan "github.com/alexchen/test/ScanService"
	"strings"
	"time"
)

// 备份缓存中的数据到本地
func main() {
	start := time.Now()
	defer fmt.Println(time.Since(start))

	var data []string
	filename := fmt.Sprintf("./log/"+scan.SavePath, time.Now().Format("060102"))
	data = Cache.SMembers("have_save_file")
	for _, v := range data {
		//scan.AppendContent(strings.TrimRight(v, "\n") + "\n")
		file.AppendContent(filename, strings.Trim(v, "\n")+"\n")
	}

	data = Cache.SMembers("laravel_database_files")
	for _, v := range data {
		file.AppendContent(filename, v+"\n")
	}
}
