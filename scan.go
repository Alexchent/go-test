package main

import scan "github.com/alexchen/test/ScanService"

func main() {
	scan.Do()

	// 缓存数据持久化到本地
	//scan.SaveCache()
}