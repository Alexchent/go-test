package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	id   = flag.Int("id", 123, "i")
	date = flag.String("date", time.Now().Format("20060102"), "d")
	open string
)

// go run flag.go --id=89 --date=20220624 --open=no
func main() {
	flag.StringVar(&open, "open", "yes", "是否开启")
	// 注意使用这些传入的参数前，需要解析命令行参数
	flag.Parse()

	fmt.Println("id:", *id)
	fmt.Println("date:", *date)
	fmt.Println(open)
}
