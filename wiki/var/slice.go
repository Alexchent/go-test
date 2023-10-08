package main

import "fmt"

func main() {
	var numbers []int //var 声明 slice，值为nil nil切片和nil相等，一般用来表示一个不存在的切片
	printSlice(numbers)
	fmt.Println(numbers == nil)
	//len=0 cap=0 slice=[]
	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)
	//len=1 cap=1 slice=[0]

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)
	//len=5 cap=6 slice=[0 2 3 4]

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
