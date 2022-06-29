package main

import "fmt"

// goto不推荐使用。以免造成程序流程混乱
func main() {
	var a int = 10

	//LOOP为标记位置
LOOP:
	for a < 20 {
		if a == 15 {
			a += 1
			goto LOOP //goto程序直接跳转到goto位置
		}
		fmt.Printf("a的值：%d\n", a)
		a++
	}
}
