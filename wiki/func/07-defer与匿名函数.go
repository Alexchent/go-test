package main

import "fmt"

func f1() (ret int) {
	defer func() {
		ret++
	}()

	return 1
}

func main() {
	fmt.Println(f1())
}

// 分析发现 函数f返回ret，其初始值为1， ret++执行是在return 1后面，因此ret变成了2
