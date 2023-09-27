package main

import "fmt"

func carPooling(trips [][]int, capacity int) bool {
	len := 10
	nums := make([]int, len)
	for _, trip := range trips {
		incr, start, end := trip[0], trip[1], trip[2]
		nums[start] += incr
		nums[end] -= incr
	}

	for i := 0; i < len; i++ {
		nums[i+1] += nums[i]
		if nums[i] > capacity {
			return false
		}
	}
	fmt.Println(nums)
	return true
}

func main() {
	trips := [][]int{{9, 0, 1}, {3, 3, 7}}
	s := carPooling(trips, 4)
	fmt.Println(s)
}
