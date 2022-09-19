package main

import "fmt"

func main() {
	mp := make(map[string]string)

	mp["china"] = "beijing"
	mp["Japan"] = "dongjing"

	/*查看元素在集合中是否存在 */
	country := "Japan"
	capital, ok := mp[country]
	if ok {
		fmt.Println(country, "的首都是", capital)
	} else {
		fmt.Println(country, "的首都不存在")
	}

	// 从map中删除元素
	delete(mp, "Japan")
	fmt.Println(mp)
}
