package main

import "fmt"

// ShiYan 当你需要将映射作为结构体字段的时候，你可以这样定义它：
type ShiYan struct {
	Name    string
	Friends map[string]*ShiYan
}

func main() {
	goku := ShiYan{
		Name:    "goku",
		Friends: make(map[string]*ShiYan),
	}
	goku.Friends["kill"] = &ShiYan{
		Name:    "alex",
		Friends: nil,
	}

	fmt.Println(goku.Friends["kill"])
}
