package main

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
