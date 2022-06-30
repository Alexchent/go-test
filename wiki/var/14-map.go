package main

import "fmt"

type data struct {
	name string
}

func main() {
	map1 := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}
	for k, v := range map1 {
		fmt.Printf("%d: %s\n", k, v)
	}

	map1[2] = "php" // 修改map
	fmt.Println(map1)

	maps := map[string]data{"x": {"tony"}}
	// maps["x"].name = "one" 注意 如果你有一个struct值的map，你无法更新单个的struct值。
	maps["x"] = data{"yvone"} // 正确的方式
	fmt.Println(maps)

	//使用指针map
	map3 := map[string]*data{"x": {"tony"}}
	map3["x"].name = "hello"
	fmt.Println(map3["x"])

	// 创建一个空map
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	lookup["hello"] = 100

	// 判断map内是否存在某个key
	value, exists := lookup["vegeta"]
	// prints 0, false
	// 0 is the default value for an integer
	fmt.Println(value, exists)

	// 我们使用 len 方法类获取映射的键的数量。使用 delete 方法来删除一个键对应的值：
	fmt.Println(len(lookup))

	delete(lookup, "goku")

	fmt.Println(lookup)
}
