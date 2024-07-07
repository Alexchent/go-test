package main

import (
	"fmt"
	myFile "github.com/alexchen/go_test/File"
	"github.com/alexchen/go_test/cache/redis"
	"github.com/alexchen/go_test/scan/scan_service"
	"os"
	"strings"
	"time"
)

// 备份缓存中的数据到本地
func main() {
	start := time.Now()
	defer fmt.Println(time.Since(start))

	//fmt.Printf("请输入要保存文件的文职:\n, 默认 $HOME/scanlog")
	//var path string
	//_, err := fmt.Scan(&path)
	//if err != nil {
	//	return
	//}
	homePath, err := os.UserHomeDir()
	dir := homePath + "/scanLog/"
	myFile.CreateDateDir(dir)

	filename := dir + fmt.Sprintf(scan.SavePath, time.Now().Format("060102"))
	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cacheKey := scan.CacheKey
	data := redis.SMembers(cacheKey)
	for _, v := range data {
		fmt.Println(v)
		fd.WriteString(strings.Trim(v, "\n"))
	}
	fmt.Println(fmt.Sprintf("总行数：%d", len(data)))
}
