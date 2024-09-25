package hs

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
