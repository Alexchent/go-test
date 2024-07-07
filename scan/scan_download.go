package main

import (
	"github.com/alexchen/go_test/scan/scan_service"
)

func main() {
	scan.Start("/Users/chentao/Downloads/xunlei")

	// 缓存数据持久化到本地
	//scan.SaveCache()
}
