package main

import (
	"fmt"
	"sort"
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

func sortMapByKey(mp map[string]int) {
	// 将map的key放到切片中
	var newMap = make([]string, 0)
	for k, _ := range mp {
		newMap = append(newMap, k)
	}

	//对切片排序
	sort.Strings(newMap)

	//3.遍历切片，然后按key来输出map的值
	for _, v := range newMap {
		fmt.Printf("根据key排序后的新集合为:[%v]=%v \n", v, mp[v])
	}
}

//根据value排序
func sortMapValue(mp map[string]int) {
	var newMp = make([]int, 0)
	var newMpKey = make([]string, 0)
	for oldk, v := range mp {
		newMp = append(newMp, v)
		newMpKey = append(newMpKey, oldk)
	}
	sort.Ints(newMp)
	for k, v := range newMp {
		fmt.Printf("根据value排序后的新集合为:[%v]=%v \n", newMpKey[k], v)
		//fmt.Println("根据value排序后的新集合为  key:", newMpKey[k], "    value:", v)
	}
}
