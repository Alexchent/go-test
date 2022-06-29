package main

import "fmt"

// 返回一个“返回int的函数”  闭包实现
func fibonacci3() func() int {
	ak := []int{0, 1}
	start := 0
	var res int
	return func() int {
		if start < 2 {
			res = ak[start]
		} else {
			res = ak[start-2] + ak[start-1]
			ak = append(ak, res)
		}
		start++
		return res
	}
}

func main() {
	f := fibonacci3()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
