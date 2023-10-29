package main

import "fmt"

func main() {
	a := 100

	// 不同于打大多数语言，go的不会自动从一个分支代码调到下一个分支执行，所以不需要break
	// 相反如果像进入下个代码块，需要加入fallthrough
	// 注意 fallthrough 只能出现在分支代码块的最后一句，且后面必须有可以分支代码块
	switch a {
	case 100:
		fmt.Println(a)
		fallthrough
	case 10, 20:
		// case 可以等于多个值
		fmt.Println("no")
	default:
		// 没有匹配的则执行default
		fmt.Println("default")
	}
}
