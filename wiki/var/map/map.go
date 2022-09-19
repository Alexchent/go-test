package main

import (
	"fmt"
)

func main() {
	// var 声明一个变量，但没有分配内存
	var countryCapital1 map[string]string
	//fmt.Printf("%v \n", countryCapital)
	fmt.Println("var countryCapital1 == nil ? ", countryCapital1 == nil)

	// 直接赋值，由程序推导变量数据类型
	countryCapital2 := map[string]string{"china": "beijing", "riben": "dongjing"}
	fmt.Println(countryCapital2)

	// make 分配内存
	countryCapital := make(map[string]string)
	//fmt.Printf("%v \n", countryCapital)
	fmt.Println("make countryCapital == nil ? ", countryCapital == nil)

	// 分配了内存的 map 才可以赋值
	countryCapital["France"] = "巴黎"
	countryCapital["Japan"] = "东京"
	countryCapital["India"] = "新德里"

	for index, value := range countryCapital {
		fmt.Println(index, "首都是", value)
		//fmt.Println(index, "首都是", countryCapital[index])
	}
}
