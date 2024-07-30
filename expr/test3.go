package main

import (
	"fmt"
	"github.com/antonmedv/expr"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	//var out any
	for i := 0; i < 1; i++ {
		program, _ := expr.Compile(`Value in 1690855200..1691632800 || Value in 1692064800..1692496800`, expr.AsBool())
		out, _ := expr.Run(program, map[string]interface{}{"Value": 1690855200})
		fmt.Println(out)

		program, _ = expr.Compile(`Value in ["monday", "tuesday", "wednesday","thursday","friday","saturday","sunday"]`, expr.AsBool())
		out, _ = expr.Run(program, map[string]interface{}{"Value": "monday"})
		fmt.Println(out)

		program, _ = expr.Compile(`Value in 1000..1200 || Value in 1400..1730`, expr.AsBool())
		out, _ = expr.Run(program, map[string]interface{}{"Value": 1100})
		fmt.Println(out)

		program, _ = expr.Compile(`Value in ["5G", "wifi"]`, expr.AsBool())
		out, _ = expr.Run(program, map[string]interface{}{"Value": "5G"})
		fmt.Println(out)

		program, _ = expr.Compile(`Value >= 1010003`, expr.AsBool())
		out, _ = expr.Run(program, map[string]interface{}{"Value": 1010003})
		fmt.Println(out)

		program, _ = expr.Compile(`Value in [1010003,1011002]`, expr.AsBool())
		out, _ = expr.Run(program, map[string]interface{}{"Value": 1010003})
		fmt.Println(out)

		program, _ = expr.Compile(`Value > 8`, expr.AsBool())
		out, _ = expr.Run(program, map[string]interface{}{"Value": 9})
		fmt.Println(out)

		program, _ = expr.Compile(`Value in ["huaw", "oppo"]`, expr.AsBool())
		out, _ = expr.Run(program, map[string]interface{}{"Value": "huaw"})
		fmt.Println(out)

		program, _ = expr.Compile(`Value contains "huaw" || Value contains "oppo"`, expr.AsBool())
		out, _ = expr.Run(program, map[string]interface{}{"Value": "hua.com"})
		fmt.Println(out)

		program, _ = expr.Compile(`Value == 1010003 && Key == 1`, expr.AsBool())
		out, _ = expr.Run(program, map[string]interface{}{"Value": 1010003, "Key": 2})
		fmt.Println(out)
	}

	//fmt.Print(output)

	fmt.Println(time.Since(start))

	//env := map[string]interface{}{
	//	"NetWork": "5G",
	//}
	//data := []string{"5G", "wifi"}
	//CheckNetWork(data, env)
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
