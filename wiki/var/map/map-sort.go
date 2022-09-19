package main

import (
	"fmt"
	"sort"
)

func main() {
	countryCapital := make(map[string]string)
	// 分配了内存的 map 才可以赋值
	countryCapital["2"] = "3巴黎"
	countryCapital["1"] = "1东京"
	countryCapital["3"] = "2新德里"

	//for k, v := range countryCapital {
	//	fmt.Println(k, v)
	//}

	//sortMapByKey(countryCapital)

	sortMapValue(countryCapital)
}

func sortMapByKey(mp map[string]string) {
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
func sortMapValue(mp map[string]string) {
	var newMp = make([]string, 0)
	var newMpKey = make([]string, 0)
	for oldk, v := range mp {
		newMp = append(newMp, v)
		newMpKey = append(newMpKey, oldk)
	}
	sort.Strings(newMp)
	for k, v := range newMp {
		fmt.Printf("根据value排序后的新集合为:[%v]=%v \n", newMpKey[k], v)
		//fmt.Println("根据value排序后的新集合为  key:", newMpKey[k], "    value:", v)
	}
}
