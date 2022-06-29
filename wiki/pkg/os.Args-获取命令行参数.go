package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 通过 os.Args 获取
	// go run os.Args-获取命令行参数.go name hello
	list := os.Args
	fmt.Println(list)
	//第一是命令行本身
	for k, v := range list {
		fmt.Println(k, v)
	}

	// 通过 flag 获取
	// go run os.Args-获取命令行参数.go -name=hello -id=13
	name := flag.String("name", "alex", "输入用户名")
	id := flag.Int("id", 2, "ID")
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*id)
}
