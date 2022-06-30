package main

import "fmt"

type Author2 struct {
	Name         string
	Publications []string
}

func main() {
	luxun := Author2{"鲁迅", []string{"闰土", "早"}}
	fmt.Printf("%v\n", luxun)  // %v 输出struct个成员的值
	fmt.Printf("%+v\n", luxun) // %+v 输出struct个成员的名称和值
	fmt.Printf("%#v\n", luxun) // %#v 输出struct名称和struct个成员的名称、类型、值
}
