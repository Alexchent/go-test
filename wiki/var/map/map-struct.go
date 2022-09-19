package main

import "fmt"

type data struct {
	name string
}

// ShiYan 当你需要将映射作为结构体字段的时候，你可以这样定义它：
type ShiYan struct {
	Name    string
	Friends map[string]*ShiYan
}

func main() {
	maps := map[string]data{"x": {"tony"}}
	// maps["x"].name = "one" 注意 如果你有一个struct值的map，你无法更新单个的struct值。
	maps["x"] = data{"yvone"} // 正确的方式
	fmt.Println(maps)

	//使用指针map
	map3 := map[string]*data{"x": {"tony"}}
	map3["x"].name = "hello"
	fmt.Println(map3["x"])

	goku := ShiYan{
		Name:    "goku",
		Friends: make(map[string]*ShiYan),
	}
	goku.Friends["kill"] = &ShiYan{
		Name:    "alex",
		Friends: nil,
	}

	fmt.Println(goku)
}
