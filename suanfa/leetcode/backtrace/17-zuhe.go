package hs

import "fmt"

// 17.电话号码组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
func letterCombinations(digits string) []string {
	var phoneMap map[string]string = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var combinations []string
	if len(digits) == 0 {
		return combinations
	}
	// 取出第一层，而后在这一层的基础上，扩张
	number := string(digits[0])
	str := phoneMap[number]
	for _, char := range str {
		combinations = append(combinations, string(char))
	}
	// 遍历数字
	for i := 1; i < len(digits); i++ {
		number = string(digits[i])
		str = phoneMap[number]
		var next []string
		// 遍历数字对应的字母
		for _, char := range str {
			// 在上次得到的组合基础上进行扩展
			for _, item := range combinations {
				next = append(next, item+string(char))
			}
		}
		combinations = next
		next = nil
	}
	return combinations
}

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。
func combine2(n int, k int) [][]int {
	var (
		res    [][]int
		subset []int
	)

	var backtrace func(int)
	backtrace = func(start int) {
		// 剪枝
		if len(subset)+(n-start+1) < k {
			return
		}
		// 将合法的组合加入结果集
		if len(subset) == k {
			res = append(res, append([]int{}, subset...))
			return
		}
		// 正式的搜索过程
		for i := start; i <= n; i++ {
			subset = append(subset, i)
			backtrace(i + 1)
			// 回退一步
			subset = subset[:len(subset)-1]
		}
	}
	// 启动回溯
	backtrace(1)
	return res
}

// permute 全排列
func permute(nums []int) [][]int {
	var res [][]int
	dfs(nums, []int{}, &res)
	return res
}

func dfs(nums, path []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i, num := range nums {
		// 深度优先遍历
		// 取出数组中的第i个放到临时存储空间path中，数组剩下的部分作为新的num。此时path中就是深度优先第一轮的到的数
		dfs(append(append([]int{}, nums[:i]...), nums[i+1:]...), append(path, num), res)
	}
}

// 数组总和
// 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
//
// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
//
// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
func combinationSum(candidates []int, target int) [][]int {
	var comb []int
	var ans [][]int
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		// 遍历到最后一个数即退出
		if idx == len(candidates) {
			return
		}
		// 剪枝，一路深度优先遍历退出
		if target == 0 {
			ans = append(ans, append([]int{}, comb...))
			return
		}
		// 直接跳过进入下一层
		dfs(target, idx+1)

		// 处理当前层
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			fmt.Println(comb)
			dfs(target-candidates[idx], idx)
			// 撤销，回到上一步
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return ans
}
