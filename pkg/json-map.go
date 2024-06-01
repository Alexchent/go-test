package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// map 转 json
	m := make(map[string]interface{})
	m["name"] = "alex"
	m["age"] = 18
	m["sex"] = "man"
	m["parents"] = []string{"tony", "yvone"}
	ms, _ := json.Marshal(m)
	fmt.Println("map 转 json ==== ", string(ms))

	// json 转 map 注意要先将json字符串转化成 []byte
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	err1 := json.Unmarshal(b, &f)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	fmt.Println("json 转 map ====", f)
}
