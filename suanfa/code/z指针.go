package code

// 双指针算在的典型场景
// 一般是在有序数组中进行操作，比如二分查找
// 也可以在有序链表中进行操作，比如删除链表中的重复元素
// 也可以在两个有序数组中进行操作，比如合并两个有序数组
// 也可以在一个有序数组中进行操作，比如删除指定元素
// 也可以在一个有序数组中进行操作，比如两数之和

func TwoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{-1, -1}
}

func RemoveDuplicates(nums []int) int {
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
