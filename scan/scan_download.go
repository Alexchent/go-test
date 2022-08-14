package main

import (
	"github.com/alexchen/go_test/scan/ScanService"
)

func main() {
	scan.Start("/Users/chentao/Downloads/xunlei")

	// 缓存数据持久化到本地
	//scan.SaveCache()
}
