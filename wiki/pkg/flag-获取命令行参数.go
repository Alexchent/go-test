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

//  go run flag-获取命令行参数.go --id=89 --date=20220624
func main() {

	flag.Parse() //解析命令行参数

	fmt.Println("id:", *id)

	fmt.Println("date:", *date)
}
