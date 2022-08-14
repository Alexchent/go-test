package main

import (
	"fmt"
	"github.com/alexchen/go_test/scan/ScanService"
)

func main() {
	var path string
	fmt.Printf("请输入要扫描的目录:\n")

	_, err := fmt.Scan(&path)
	if err != nil {
		return
	}
	//scan.Start(path)
	path = "/Users/chentao/Documents/转正"
	scan.Do(path)
	// 缓存数据持久化到本地
	//scan.SaveCache()
}
