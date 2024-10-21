package hs

import "fmt"

// 括号生成
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
func generateParenthesis(n int) []string {
	var ans []string
	var backtrace func(cur []rune, left, right int)
	// 通过跟踪已生成的括号字符串cur、已使用的左括号数量和右括号数量来生成所有可能的组合
	backtrace = func(cur []rune, left, right int) {
		fmt.Println(string(cur), left, right)
		// 剪枝条件，当生成的字符串长度达到2*n时，说明找到了一个有效的括号组合，将其添加到结果列表中
		if len(cur) == 2*n {
			ans = append(ans, string(cur))
			return
		}
		// 如果左括号的数量小于n，则可以添加一个左括号继续递归
		if left < n {
			cur = append(cur, '(')
			backtrace(cur, left+1, right)
			cur = cur[:len(cur)-1]
		}
		// 如果右括号的数量小于左括号的数量，则可以添加一个右括号继续递归。
		if right < left {
			cur = append(cur, ')')
			backtrace(cur, left, right+1)
			cur = cur[:len(cur)-1]
		}
	}

	backtrace([]rune{}, 0, 0)
	return ans
}
