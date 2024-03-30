package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

// https://github.com/spf13/pflag
// go run pflag.go --name=12 --flagvar=22 -v=false
func main() {
	// pflag 支持4种命令行参数定义方式
	// 1. 支持长选项、默认值和使用文本，并将标志的值存储在指针中。
	var _ = flag.String("name", "colin", "Input Your Name")
	// 2. 支持长选项、短选项、默认值和使用文本，并将标志的值存储在指针中。
	var _ = flag.StringP("name", "n", "colin", "Input Your Name")
	// 3. 绑定变量
	var flagVar int
	flag.IntVar(&flagVar, "flagVar", 1234, "help message for flagname")
	// 4. 绑定变量，同时支持短选项
	var version bool
	flag.BoolVarP(&version, "version", "v", true, "print version")

	flag.Parse() //定义好输入后，调用flag.Parse()，进行数据解析

	fmt.Println("flagVar", flagVar)
	fmt.Println("version", version)
}
