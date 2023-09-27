package code

import "math"

// BinaryFind 判断一个数字是否在有序列表中
// 二分查找
// 时间复杂度：O(logn)
func BinaryFind(a int, arr []int) bool {
	left, right := 0, len(arr)
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == a {
			return true
		} else if a < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

//给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。

func MajorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			return v
		}
	}
	return -1
}

// 注意：go 代码由 chatGPT🤖 根据我的 cpp 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码不保证正确性，仅供参考。如有疑惑，可以参照我写的 cpp 代码对比查看。

func minWindow(s string, t string) string {
	need := make(map[byte]int)   // 用于统计需要凑齐的字符
	window := make(map[byte]int) // 记录滑动窗口内已有字符的个数
	for i := range t {
		need[t[i]]++
	}
	left, right := 0, 0               // 滑动窗口的左右边界
	valid := 0                        // 判断窗口中是否已经包含了字串 t 中所有字符
	start, length := 0, math.MaxInt32 // 最小覆盖子串的起始索引及长度
	for right < len(s) {              // 当 right 小于 s 的长度时，继续循环
		c := s[right] // c 是将要加入窗口中的字符
		right++
		if _, ok := need[c]; ok { // 如果这个字符在字串 t 中需要的话
			window[c]++               // 加入窗口中
			if window[c] == need[c] { // 如果字符 c 在窗口中的数量已经满足其在字串 t 中的数量
				valid++ // 计数器 valid 加一
			}
		}
		for valid == len(need) { // 如果滑动窗口中的字符已经完全覆盖字串 t 中的字符
			if right-left < length { // 如果此时的覆盖子串更短
				start = left          // 更新最小覆盖子串的起始索引
				length = right - left // 更新最小子串的长度
			}
			d := s[left]              // d 是将要移出窗口的字符
			left++                    // 左侧窗口右移
			if _, ok := need[d]; ok { // 如果这个字符在字串 t 中需要的话
				if window[d] == need[d] { // 如果这个字符已经满足了他在字串 t 中的需求
					valid-- // 计数器 valid 减一
				}
				window[d]-- // 移出窗口
			}
		}
	}
	if length == math.MaxInt32 { // 如果最小子串长度没有更新，则返回空格
		return ""
	}
	return s[start : start+length] // 返回最小覆盖子串
}
