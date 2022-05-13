package main

import (
	"fmt"
	scan "github.com/alexchen/test/ScanService"
)

func main() {
	var path string
	fmt.Printf("请输入要扫描的目录:\n")

	_, err := fmt.Scan(&path)
	if err != nil {
		return
	}
	scan.Start(path)

	// 缓存数据持久化到本地
	//scan.SaveCache()
}
