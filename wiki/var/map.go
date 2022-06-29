package main

import "fmt"

func main() {
	// 创建集合
	var countryCapital map [string] string
    
    countryCapital = make(map[string]string)

    countryCapital["France"] = "巴黎"
    countryCapital["Japan"] = "东京"
    countryCapital["India"] = "新德里"

    for index,_ := range countryCapital {
    	//fmt.Println(index,"首都是",value)
    	fmt.Println(index,"首都是",countryCapital[index])
    }

    /*查看元素在集合中是否存在 */
    country := "American"
    capital, ok := countryCapital[ country ] /*如果确定是真实的,则存在,否则不存在 */
    if (ok) {
        fmt.Println(capital, "的首都是", ok)
    } else {
        fmt.Println(country, "的首都不存在")
    }

    delete(countryCapital,"Japan")

    fmt.Println(countryCapital)
}