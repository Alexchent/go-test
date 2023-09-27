package code

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
func Rotate(nums []int, k int) {
	k = k % len(nums)
	// 三次反转
	// 1. 先反转整个数组
	Reverse(nums, 0, len(nums)-1)
	// 2. 反转前k个元素
	Reverse(nums, 0, k-1)
	// 3. 反转后面的元素
	Reverse(nums, k, len(nums)-1)
}
