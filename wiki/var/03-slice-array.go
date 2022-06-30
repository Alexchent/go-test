package main

import "fmt"

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice1 := a[2:4:8]
	fmt.Println(slice1) // [2, 3]

	slice2 := slice1[1:6] // 注意切片长度不能越过容量
	fmt.Println(slice2)   // [3, 4, 5, 6, 7]

	testSlice()

	testMap()
}

func testSlice() {
	x := []int{1, 2, 3}

	func(arr []int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)
	fmt.Println("slice类型的变量发生改变")
	fmt.Println(x)
}

func testMap() {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)
	fmt.Println("array类型的变量没有改变")
	fmt.Println(x)
}
