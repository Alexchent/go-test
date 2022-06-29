package main

import (
	"fmt"
)

func main() {
	//var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	a := []int{2, 3, 1, 6, 8, 7, 4, 5}
	SortPop(&a)
	fmt.Println(a)
}

// SortPop 冒泡排序 按引用传递
func SortPop(arr2 *[]int) {
	arr := *arr2
	num := len(arr)
	for i := 0; i < num-1; i++ {
		for j := 0; j < num-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
