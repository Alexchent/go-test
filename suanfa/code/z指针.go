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

// IsSubsequence 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
// 输入：s = "abc", t = "ahbgdc"
// 输出：true
func IsSubsequence(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

// 大小转小写
func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

// leetcode 有效的括号
func isValid(s string) bool {
	// 奇数直接返回false
	if len(s)%2 == 1 {
		return false
	}
	// 用map存储括号
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	// 用栈存储括号
	stack := []byte{}
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		// 如果是右括号
		if pairs[s[i]] > 0 {
			// 如果栈为空或者栈顶元素不是对应的左括号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，则压入栈
			stack = append(stack, s[i])
		}
	}
	// 如果栈为空，则返回true
	return len(stack) == 0
}

// 长度最小的子数组
// 给定一个含有n个正整数的数组和一个正整数target
// 找出该数组中满足其和 ≥ target 的长度最小的连续子数组
// 如果不存在符合条件的子数组，则返回 0
func minSubArrayLen(target int, nums []int) int {
	// 滑动窗口
	left, right := 0, 0
	sum := 0
	minLen := len(nums) + 1
	for right < len(nums) {
		// 窗口右移
		sum += nums[right]
		right++
		// 窗口左移
		for sum >= target {
			if right-left < minLen {
				minLen = right - left
			}
			sum -= nums[left]
			left++
		}
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}

// 无重复字符的最长子串
// 给定一个字符串，请你找出其中不含有重复字符的最长子串的长度
func lengthOfLongestSubstring(s string) int {
	// 滑动窗口
	left, right := 0, 0
	maxLen := 0
	// 用map存储字符出现的次数
	m := map[byte]int{}
	for right < len(s) {
		// 窗口右移
		m[s[right]]++
		right++
		// 窗口左移
		for m[s[right-1]] > 1 {
			m[s[left]]--
			left++
		}
		maxLen = max(maxLen, right-left)
	}
	return maxLen
}

// 同构字符串
// 给定两个字符串s和t，判断它们是否是同构的
// 如果s中的字符可以被替换得到t，则两个字符串是同构的
// 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序
// 两个字符不能映射到同一个字符上，但字符可以映射自己本身
func isIsomorphic(s, t string) bool {
	// 用map存储字符出现的次数
	m1, m2 := map[byte]byte{}, map[byte]byte{}
	for i := 0; i < len(s); i++ {
		if m1[s[i]] == 0 {
			m1[s[i]] = t[i]
		}
		if m2[t[i]] == 0 {
			m2[t[i]] = s[i]
		}
		if m1[s[i]] != t[i] || m2[t[i]] != s[i] {
			return false
		}
	}
	return true
}
