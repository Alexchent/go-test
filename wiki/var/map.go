package main

import "fmt"

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

	/*查看元素在集合中是否存在 */
	country := "American"
	capital, ok := countryCapital[country] /*如果确定是真实的,则存在,否则不存在 */
	if ok {
		fmt.Println(capital, "的首都是", ok)
	} else {
		fmt.Println(country, "的首都不存在")
	}

	delete(countryCapital, "Japan")

	fmt.Println(countryCapital)
}
