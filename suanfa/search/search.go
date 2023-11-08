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
