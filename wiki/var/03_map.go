package main

import "fmt"

func main() {

	//var city map[string]int
	//city["apple"] = 1
	// 在一个nil的mao中添加元素是会生成一个运行时的panic

	// 1. 先申明map变量，再赋值
	var country = make(map[string]string)
	country["china"] = "beijing"
	country["japan"] = "dongjing"
	fmt.Println(country)

	// 2. 声明的同时赋值
	countryCapitalMap := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
		"India":  "New delhi",
	}

	/* 打印地图 */
	for country, cap := range countryCapitalMap {
		//fmt.Println(country, "首都是", countryCapitalMap[country])
		// 注意 range 迭代的顺序是随机的
		fmt.Println(country, "首都是", cap)
	}

	// 删除元素
	delete(countryCapitalMap, "France")

	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}
