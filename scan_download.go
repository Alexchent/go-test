package main

import (
	scan "github.com/alexchen/test/ScanService"
)

func main() {
	scan.Start("/Users/chentao/Downloads/xunlei")

	// 缓存数据持久化到本地
	//scan.SaveCache()
}
