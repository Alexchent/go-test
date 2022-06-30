package main

import "fmt"

func main() {
	// 数组是比较特殊的，var 声明时已经带上长度，值永远不可能是nil, 因此可以通过 a[index]=value 的方式赋值
	var a [2]int
	var b []int

	a[0] = 10 // 数组是比较特殊的
	//b[0] = 10 // panic: runtime error: index out of range [0] with length 0
	b = append(b, 10, 20)
	fmt.Println(a, b)

	testSlice()
	testMap()
}

// 切片作为函数的参数，传的是指针
func testSlice() {
	x := []int{1, 2, 3}

	func(arr []int) {
		arr[1] = 7
		fmt.Println(arr)
	}(x)
	fmt.Println("slice类型的变量 —— 改变")
	fmt.Println(x)
}

// 切片作为函数的参数，传的是值
func testMap() {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[1] = 7
		fmt.Println(arr)
	}(x)
	fmt.Println("array类型的变量 —— 不变")
	fmt.Println(x)
}
