package code

import "fmt"

// 跳跃游戏
// 给定一个非负整数数组，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个位置。
func canJump(nums []int) bool {
	// 从后往前贪心
	lastPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		// 如果当前位置能够到达最后一个位置
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}

// 跳跃游戏 II
// 给定一个非负整数数组，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
func Jump(nums []int) int {
	// 贪心算法
	// 从前往后
	// 每次在可跳范围内选择可以使得跳的更远的位置
	end := 0    // 上次跳跃可达范围右边界
	maxPos := 0 // 当前能跳到的最远位置
	step := 0   // 跳跃次数
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, i+nums[i]) // 更新当前能跳到的最远位置
		if i == end {                   // 遇到上次跳跃可达范围右边界，更新边界并将跳跃次数+1
			end = maxPos // 更新边界
			step++
		}
		fmt.Println(end, maxPos, step)
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
