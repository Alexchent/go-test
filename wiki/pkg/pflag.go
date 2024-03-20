package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

// https://github.com/spf13/pflag
// go run pflag.go --flagname=12 --flagvar=22
func main() {

	var ip *int = flag.Int("flagname", 1234, "help message for flagname")

	var flagvar int
	flag.IntVar(&flagvar, "flagvar", 1234, "help message for flagname")

	flag.Parse() //定义好输入后，调用flag.Parse()，进行数据解析

	fmt.Println(*ip)
	fmt.Println(flagvar)

}
