package main

import "fmt"

// go的内存分配有两种，栈和堆，栈的操作高效但空间有限，一般函数的参数、局部变量、函数的调用帧会分配到栈上
// 内存逃逸，即原本应该分配到栈上的数据，被分配到了堆上
// 内存逃逸，主要有3种情况，
//     1. 变量类型不确定，比如用interface{}作为数据类型
//     2. 变量在函数外部存在引用
//     3. 变量所占内存较大

// 进行逃逸分析 go build -gcflags '-m -l' cgflags.go
func main() {
	a := 666
	fmt.Println(a) // Println的参数是any，所以第一种情况
}

// 返回指针，第2种情况
func foo() *int {
	a := 666
	return &a
}

func foo2() {
	// 局部变量较大
	s := make([]int, 10000, 10000)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
}
