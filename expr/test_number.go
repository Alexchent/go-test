package main

import (
	"fmt"
	"github.com/antonmedv/expr"
	"strings"
)

func main() {
	// 范围
	program, _ := expr.Compile(`Value in 1000..1200 || Value in 1400..1730`, expr.AsBool())
	out, _ := expr.Run(program, map[string]interface{}{"Value": 1100})
	fmt.Println(out)

	program, _ = expr.Compile(`Value in [1,10]`, expr.AsBool())
	out, _ = expr.Run(program, map[string]interface{}{"Value": 1})
	fmt.Println(out)
	out, _ = expr.Run(program, map[string]interface{}{"Value": 2})
	fmt.Println(out)

	program, _ = expr.Compile(`Value == 8.8`, expr.AsBool())
	out, _ = expr.Run(program, map[string]interface{}{"Value": 8.8})
	fmt.Println(out)

	program, _ = expr.Compile(`Value != 8`, expr.AsBool())
	out, _ = expr.Run(program, map[string]interface{}{"Value": 8})
	fmt.Println(out)

	program, _ = expr.Compile(`Value >= 100`, expr.AsBool())
	out, _ = expr.Run(program, map[string]interface{}{"Value": 101})
	fmt.Println(out)
	out, _ = expr.Run(program, map[string]interface{}{"Value": 99})
	fmt.Println(out)

	program, _ = expr.Compile(`Value < 8`, expr.AsBool())
	out, _ = expr.Run(program, map[string]interface{}{"Value": 7})
	fmt.Println(out)

	env := map[string]interface{}{
		"NetWork": "5G",
	}
	data := []string{"5G", "wifi"}
	CheckNetWork(data, env)
}

func CheckNetWork(data []string, env map[string]interface{}) {
	joinedData := fmt.Sprintf(`["%s"]`, strings.Join(data, `", "`))
	code := fmt.Sprintf("NetWork in %s", joinedData)
	program, err := expr.Compile(code, expr.AsBool())
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	fmt.Print(output)
}
