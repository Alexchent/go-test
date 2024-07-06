package search

// BinarySearch 二分查找
// 二分查找的前提是数组有序
// 二分查找的时间复杂度是 O(logn)
// 思路：取数组的中间值，如果中间值等于目标值，则返回中间值的索引
// 如果中间值大于目标值，则在左边继续查找
// 如果中间值小于目标值，则在右边继续查找
// 不断缩小区间，直到找到目标值，或者区间缩小为0
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		min := (left + right) / 2
		if arr[min] == target {
			return min
		}
		if arr[min] > target {
			right = min - 1
		} else {
			left = min + 1
		}
	}
	return -1
}

// 二分查找
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		//mid := left + (right-left)/2
		mid := (left + right) / 2
		switch {
		case arr[mid] == target:
			return mid
			// 如果中间值大于目标值，则在左边查找
		case arr[mid] > target:
			right = mid - 1
			// 如果中间值小于目标值，则在右边查找
		case arr[mid] < target:
			left = mid + 1
		}
	}
	return -1
}

// 一个数组。一个目标数，找两个数之和是目标数的下标
func twoSum(nums []int, target int) []int {
	// 用map存储数组元素和下标
	numMap := make(map[int]int)
	// 遍历数组
	for i, num := range nums {
		// 如果map中存在目标值减去当前值的key，则返回
		if j, ok := numMap[target-num]; ok {
			return []int{j, i}
		}
		// 否则将当前值存入map，这样
		numMap[num] = i
	}
	return nil
}
