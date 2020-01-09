package main

import "fmt"

func main() {
	// 创建集合
	var countryCapital map [string] string
    
    countryCapital = make(map[string]string)

    countryCapital["France"] = "巴黎"
    countryCapital["Japan"] = "东京"
    countryCapital["India"] = "新德里"

    for country,captial := range countryCapital {
    	fmt.Println(country,"首都是",captial)
    	fmt.Println(country,"首都是",countryCapital[country])
    }

    /*查看元素在集合中是否存在 */
    capital, ok := countryCapital[ "American" ] /*如果确定是真实的,则存在,否则不存在 */
    /*fmt.Println(capital) */
    /*fmt.Println(ok) */
    if (ok) {
        fmt.Println("American 的首都是", capital)
    } else {
        fmt.Println("American 的首都不存在")
    }

    delete(countryCapital,"Japan")


    fmt.Println(countryCapital)
}