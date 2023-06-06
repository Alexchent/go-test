package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	id   = flag.Int("id", 123, "Name to greet")
	date = flag.String("date", time.Now().Format("20060102"), "时间")
)

// go run flag-获取命令行参数.go --id=89 --date=20220624 --name=王多鱼
func main() {
	var name string
	flag.StringVar(&name, "name", "我是青蒿", "姓名")
	flag.Parse() //解析命令行参数

	fmt.Println("id:", *id)

	fmt.Println("date:", *date)

	fmt.Println(name)
}
