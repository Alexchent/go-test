package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Hobby string
}

func main() {
	p := Person{"5lmh.com", "女"}
	// 编码json
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err ", err)
	}

	// 格式化输出
	b, err = json.MarshalIndent(p, "", "	")
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))
}
