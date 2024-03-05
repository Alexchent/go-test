package main

import (
	"fmt"
	"reflect"
)

type Author struct {
	Name         string   `json:Name`
	Publications []string `json:Publication,omitempty`
}

func main() {
	// 变量struct的tag
	t := reflect.TypeOf(Author{})
	for i := 0; i < t.NumField(); i++ {

		fmt.Println(t.Field(i).Tag)
		//name := t.Field(i).Name
		//s, _ := t.FieldByName(name)
		//fmt.Println(s.Tag)
	}
	a := []string{"php", "java"}
	b := []string{"php", "java"}

	// 判断 slice、map、struct 是相等
	fmt.Println(reflect.DeepEqual(a, b))

	luxun := Author{"鲁迅", []string{"闰土", "早"}}
	fmt.Printf("%v\n", luxun)  // %v 输出struct个成员的值
	fmt.Printf("%+v\n", luxun) // %+v 输出struct个成员的名称和值
	fmt.Printf("%#v\n", luxun) // %#v 输出struct名称和struct个成员的名称、类型、值
}
