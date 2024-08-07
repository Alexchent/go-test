package main

import (
	"fmt"
	"github.com/antonmedv/expr"
)

// https://expr-lang.org/docs/language-definition

func main() {

	program, _ := expr.Compile(`Value == "5G"`, expr.AsBool())
	out, _ := expr.Run(program, map[string]interface{}{"Value": "5G"})
	fmt.Println(out)

	program, _ = expr.Compile(`Value != "5G"`, expr.AsBool())
	out, _ = expr.Run(program, map[string]interface{}{"Value": "5G"})
	fmt.Println(out)

	program, _ = expr.Compile(`Value in ["5G", "wifi"]`, expr.AsBool())
	out, _ = expr.Run(program, map[string]interface{}{"Value": "5G"})
	fmt.Println(out)

	// 字符串操作符 contains
	program, _ = expr.Compile(`Value contains "hua"`, expr.AsBool())
	out, _ = expr.Run(program, map[string]interface{}{"Value": "hua.com"})
	fmt.Println(out)

	// 字符串操作符 startWith
	program, _ = expr.Compile(`Value startsWith "hua"`, expr.AsBool())
	out, _ = expr.Run(program, map[string]interface{}{"Value": "hua.com"})
	fmt.Println(out)

	// 字符串操作符 endsWith
	program, _ = expr.Compile(`Value endsWith "com"`, expr.AsBool())
	out, _ = expr.Run(program, map[string]interface{}{"Value": "hua.com"})
	fmt.Println(out)
}
