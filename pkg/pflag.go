package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

// https://github.com/spf13/pflag
// go run pflag.go --name=12 --flagvar=22 -v=false
// 输出帮助文档 go run pflag.go -h
func main() {
	// pflag 支持4种命令行参数定义方式
	// 1. 支持长选项、默认值和使用文本，并将标志的值存储在指针中。
	var name = flag.String("name", "colin", "Input Your Name")
	// 2. 支持长选项、短选项、默认值和使用文本，并将标志的值存储在指针中。
	var nicknmae = flag.StringP("nickname", "n", "colin", "Input Your Name")
	// 3. 绑定变量
	var flagVar int
	flag.IntVar(&flagVar, "flagVar", 1234, "help message for flagname")
	// 4. 绑定变量，同时支持短选项
	var version bool
	flag.BoolVarP(&version, "version", "v", true, "print version")

	// 弃用标志或者标志的简写
	flag.CommandLine.MarkDeprecated("nickname", "please use --name instead")
	// 仅弃用标志的简写
	flag.CommandLine.MarkShorthandDeprecated("version", "please use --version only")

	flag.Parse() //定义好输入后，调用flag.Parse()，进行数据解析

	fmt.Println("name", *name)
	fmt.Println("nickname", *nicknmae)
	fmt.Println("flagVar", flagVar)
	fmt.Println("version", version)
}
